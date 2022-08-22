package motor

import (
	"context"
	"fmt"
	"net/http"

	mt "github.com/sonr-io/sonr/pkg/motor/types"
	st "github.com/sonr-io/sonr/x/schema/types"
)

func (mtr *motorNodeImpl) QueryWhatIs(ctx context.Context, request mt.QueryWhatIsRequest) (mt.QueryWhatIsResponse, error) {
	resp, err := mtr.schemaQueryClient.WhatIs(ctx, &st.QueryWhatIsRequest{
		Creator: request.Creator,
		Did:     request.Did,
	})
	if err != nil {
		return mt.QueryWhatIsResponse{}, err
	}

	// store reference to schema
	_, err = mtr.Resources.StoreWhatIs(resp.WhatIs)
	if err != nil {
		return mt.QueryWhatIsResponse{}, fmt.Errorf("store WhatIs: %s", err)
	}

	return mt.QueryWhatIsResponse{
		WhatIs: resp.WhatIs,
	}, nil
}

func (mtr *motorNodeImpl) QueryWhatIsByCreator(ctx context.Context, request mt.QueryWhatIsRequest) (mt.QueryWhatIsByCreatorResponse, error) {
	resp, err := mtr.schemaQueryClient.WhatIsByCreator(ctx, &st.QueryWhatIsCreatorRequest{
		Creator: request.Creator,
	})
	if err != nil {
		return mt.QueryWhatIsByCreatorResponse{}, err
	}

	// store reference to schema
	for _, wi := range resp.WhatIs {
		_, err = mtr.Resources.StoreWhatIs(wi)
		if err != nil {
			return mt.QueryWhatIsByCreatorResponse{}, fmt.Errorf("store WhatIs: %s", err)
		}
	}

	return mt.QueryWhatIsByCreatorResponse{
		Code:   http.StatusAccepted,
		WhatIs: resp.WhatIs,
	}, nil
}
