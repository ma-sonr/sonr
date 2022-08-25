package utils

import (
	"bytes"
	"encoding/gob"
	"time"

	"github.com/pkg/errors"
	"github.com/sonr-io/keyring"
	"github.com/sonr-io/sonr/pkg/crypto/mpc"
	"github.com/sonr-io/sonr/pkg/fs"
	mt "github.com/sonr-io/sonr/third_party/types/motor"
)

const K_SERVICE_NAME = "sonr-dev"
const K_AUTH_LIST_KEY = "auth_list"

type UserAuthList struct {
	Auths map[string]UserAuth
}

func (l UserAuthList) Add(addr string, ua UserAuth) {
	if l.Auths == nil {
		l.Auths = make(map[string]UserAuth)
	}
	l.Auths[addr] = ua
}

func (l UserAuthList) Get(addr string) (UserAuth, error) {
	ua, ok := l.Auths[addr]
	if !ok {
		return UserAuth{}, errors.New("UserAuth not found")
	}
	return ua, nil
}

func (l UserAuthList) Serialize() ([]byte, error) {
	var b bytes.Buffer
	e := gob.NewEncoder(&b)
	if err := e.Encode(l); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func DeserializeUserAuthList(b []byte) (UserAuthList, error) {
	var l UserAuthList
	bz := bytes.NewBuffer(b)
	d := gob.NewDecoder(bz)
	if err := d.Decode(&l); err != nil {
		return UserAuthList{}, err
	}
	return l, nil
}

type UserAuth struct {
	Password  string
	AesDSCKey []byte
	AesPSKKey []byte
	CreatedAt string
}

func (i UserAuth) Validate() bool {
	if len(i.Password) < 12 {
		return false
	}
	if i.AesDSCKey == nil {
		return false
	}
	if len(i.AesDSCKey) < 32 {
		return false
	}
	return true
}

func (i UserAuth) GenAccountCreateRequest() (*mt.CreateAccountRequest, error) {
	if i.Validate() {
		return &mt.CreateAccountRequest{
			Password:  i.Password,
			AesDscKey: i.AesDSCKey,
		}, nil
	}
	return nil, errors.New("Invalid User Auth Object")
}

func NewUserAuth(pwd string) (UserAuth, error) {
	aesKey, err := mpc.NewAesKey()
	if err != nil {
		return UserAuth{}, err
	}
	return UserAuth{
		Password:  pwd,
		AesDSCKey: aesKey,
		CreatedAt: time.Now().Format(time.RFC3339),
	}, nil
}

func (i UserAuth) StoreAuth(addr string, psk []byte) error {
	kc, err := fetchKCService()
	if err != nil {
		return errors.Wrap(err, "Failed to initialize keychain service")
	}
	if !i.Validate() {
		return errors.New("Invalid UserAuth Object")
	}
	i.AesPSKKey = psk

	var al UserAuthList
	al, err = GetUserAuthList()
	if err != nil {
		al = newAuthList()
	}
	al.Add(addr, i)

	bz, err := al.Serialize()
	if err != nil {
		return errors.Wrap(err, "Failed to serialize UserAuthList")
	}
	err = kc.Set(keyring.Item{
		Key:  K_AUTH_LIST_KEY,
		Data: bz,
	})
	if err != nil {
		return err
	}
	return nil
}

func HasUserAuth() bool {
	ual, err := GetUserAuthList()
	if err != nil {
		return false
	}
	return len(ual.Auths) > 0
}

func GetUserAuthList() (UserAuthList, error) {
	kc, err := fetchKCService()
	if err != nil {
		return UserAuthList{}, errors.Wrap(err, "Failed to initialize keychain service")
	}
	i, err := kc.Get(K_AUTH_LIST_KEY)
	if err != nil {
		return UserAuthList{}, err
	}
	if i.Data == nil || len(i.Data) == 0 {
		return UserAuthList{}, errors.New("Keychain Item data is invalid (empty or nil)")
	}
	return DeserializeUserAuthList(i.Data)
}

func fetchKCService() (keyring.Keyring, error) {
	return keyring.Open(keyring.Config{
		ServiceName:              K_SERVICE_NAME,
		KeychainTrustApplication: true,
		AllowedBackends: []keyring.BackendType{
			keyring.KeychainBackend,
		},
	})
}

func getSecureFolderPath() string {
	folder, err := fs.Support.CreateFolder("io.sonr.blockchain")
	if err != nil {
		return ""
	}
	return folder.Path()
}

func newAuthList() UserAuthList {
	return UserAuthList{
		Auths: make(map[string]UserAuth),
	}
}
