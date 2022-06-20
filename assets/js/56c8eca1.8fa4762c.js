"use strict";(self.webpackChunksonr_docs=self.webpackChunksonr_docs||[]).push([[5528],{3905:function(e,t,n){n.d(t,{Zo:function(){return p},kt:function(){return h}});var r=n(7294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function o(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function l(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},i=Object.keys(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var s=r.createContext({}),c=function(e){var t=r.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):o(o({},t),e)),n},p=function(e){var t=c(e.components);return r.createElement(s.Provider,{value:t},e.children)},d={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},u=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,i=e.originalType,s=e.parentName,p=l(e,["components","mdxType","originalType","parentName"]),u=c(n),h=a,m=u["".concat(s,".").concat(h)]||u[h]||d[h]||i;return n?r.createElement(m,o(o({ref:t},p),{},{components:n})):r.createElement(m,o({ref:t},p))}));function h(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var i=n.length,o=new Array(i);o[0]=u;var l={};for(var s in t)hasOwnProperty.call(t,s)&&(l[s]=t[s]);l.originalType=e,l.mdxType="string"==typeof e?e:a,o[1]=l;for(var c=2;c<i;c++)o[c]=n[c];return r.createElement.apply(null,o)}return r.createElement.apply(null,n)}u.displayName="MDXCreateElement"},1279:function(e,t,n){n.r(t),n.d(t,{assets:function(){return p},contentTitle:function(){return s},default:function(){return h},frontMatter:function(){return l},metadata:function(){return c},toc:function(){return d}});var r=n(7462),a=n(3366),i=(n(7294),n(3905)),o=["components"],l={title:"Using the CLI",id:"using-cli"},s="Overview",c={unversionedId:"run-nodes/highway/using-cli",id:"run-nodes/highway/using-cli",title:"Using the CLI",description:"The highway is a single binary which allows for interfacing with the Sonr Blockchain (see 'Using the CLI' for information on commands). The highway is also equipped with a REST server.  The following is a diagram outlining the topology of highway and available features.",source:"@site/posts/run-nodes/highway/cli.md",sourceDirName:"run-nodes/highway",slug:"/run-nodes/highway/using-cli",permalink:"/posts/run-nodes/highway/using-cli",editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/posts/run-nodes/highway/cli.md",tags:[],version:"current",frontMatter:{title:"Using the CLI",id:"using-cli"}},p={},d=[{value:"Creating a Highway Node",id:"creating-a-highway-node",level:3},{value:"Decentralized Identifiers (DIDs) Usage:",id:"decentralized-identifiers-dids-usage",level:3}],u={toc:d};function h(e){var t=e.components,n=(0,a.Z)(e,o);return(0,i.kt)("wrapper",(0,r.Z)({},u,n,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"overview"},"Overview"),(0,i.kt)("p",null,"The highway is a single binary which allows for interfacing with the Sonr Blockchain (see ",(0,i.kt)("strong",{parentName:"p"},"'Using the CLI'")," for information on commands). The highway is also equipped with a REST server.  The following is a diagram outlining the topology of highway and available features."),(0,i.kt)("p",null,"We believe the best way to onboard the next billion users is to create a cohesive end-to-end platform that\u2019s composable and interoperable with all existing protocols. For this, we built our Networking layer in Libp2p and our Layer 1 Blockchain with Starport. Our network comprises of two separate nodes: Highway and Motor, which each have a specific use case on the network. In order to maximize the onboarding experience, we developed our own Wallet which has value right out of the gate!"),(0,i.kt)("h1",{id:"using-the-cli"},"Using the CLI"),(0,i.kt)("p",null,"our Highway-sdk comes with a set of CLI commands"," "),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-none"},"serve - Serves our GRPC and HTTP servers on the specified ports in our enviorment files\n")),(0,i.kt)("h1",{id:"using-rest"},"Using REST"),(0,i.kt)("p",null,"The highway is capable of running an http server (REST) with 'serve' ports can be specified"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("p",{parentName:"li"},"Register")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("p",{parentName:"li"},"Authentication")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("p",{parentName:"li"},"Objects"," "),(0,i.kt)("ul",{parentName:"li"},(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("p",{parentName:"li"},"Create")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("p",{parentName:"li"},"Update")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("p",{parentName:"li"},"Deactivate")))),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("p",{parentName:"li"},"Buckets"),(0,i.kt)("ul",{parentName:"li"},(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("p",{parentName:"li"},"Create")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("p",{parentName:"li"},"Update")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("p",{parentName:"li"},"Deactivate")))),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("p",{parentName:"li"},"Channels"),(0,i.kt)("ul",{parentName:"li"},(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("p",{parentName:"li"},"Create")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("p",{parentName:"li"},"Hide")))),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("p",{parentName:"li"},"registry"),(0,i.kt)("ul",{parentName:"li"},(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("p",{parentName:"li"},"query")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("p",{parentName:"li"},"exists"))))),(0,i.kt)("h1",{id:"using-golang"},"Using Golang"),(0,i.kt)("p",null,"The highway node is a relayer node that helps motors interact with the sonr network. It is responsible for routing messages between motors and other relayers. The highway node\nalso provides an interface for developers to deploy custom services on the network. To have a custom build of the highway node, execute the following command on your machine:"),(0,i.kt)("h3",{id:"creating-a-highway-node"},"Creating a Highway Node"),(0,i.kt)("ol",null,(0,i.kt)("li",{parentName:"ol"},(0,i.kt)("p",{parentName:"li"},(0,i.kt)("inlineCode",{parentName:"p"},"go get -u github.com/sonr-io/sonr"))),(0,i.kt)("li",{parentName:"ol"},(0,i.kt)("p",{parentName:"li"},"Create a simple highway node with the following:"))),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},'import (\n  "github.com/sonr-io/sonr/pkg/highway"\n  "github.com/sonr-io/sonr/internal/host"\n)\n\nfunc main() {\n    // Create the node.\n    n, err := highway.NewHighway(ctx, host.WithPort(8084), host.WithWebAuthn("Sonr", "localhost", "http://localhost:8080", true))\n    if err != nil {\n        panic(err)\n    }\n}\n')),(0,i.kt)("h3",{id:"decentralized-identifiers-dids-usage"},"Decentralized Identifiers (DIDs) Usage:"),(0,i.kt)("blockquote",null,(0,i.kt)("p",{parentName:"blockquote"},"A library to parse and generate W3C ",(0,i.kt)("a",{parentName:"p",href:"https://www.w3.org/TR/did-core/"},"DID Documents")," and W3C ",(0,i.kt)("a",{parentName:"p",href:"https://www.w3.org/TR/vc-data-model/"},"Verifiable Credentials"),".")),(0,i.kt)("p",null,"Creation of a simple DID Document which is its own controller and contains an AssertionMethod."),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},'didID, err := did.ParseDID("did:snr:123")\n\n// Empty did document:\ndoc := &did.Document{\n    Context:            []did.URI{did.DIDContextV1URI()},\n    ID:                 *didID,\n}\n\n// Add an assertionMethod\nkeyID, _ =: did.ParseDIDURL("did:snr:123#key-1")\n\nkeyPair, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)\nverificationMethod, err := did.NewVerificationMethod(*keyID, did.JsonWebKey2020, did.DID{}, keyPair.Public())\n\n// This adds the method to the VerificationMethod list and stores a reference to the assertion list\ndoc.AddAssertionMethod(verificationMethod)\n\ndidJson, _ := json.MarshalIndent(doc, "", "  ")\nfmt.Println(string(didJson))\n\n// Unmarshalling of a json did document:\nparsedDIDDoc := did.Document{}\nerr = json.Unmarshal(didJson, &parsedDIDDoc)\n\n// It can return the key in the convenient lestrrat-go/jwx JWK\nparsedDIDDoc.AssertionMethod[0].JWK()\n\n// Or return a native crypto.PublicKey\nparsedDIDDoc.AssertionMethod[0].PublicKey()\n\n')),(0,i.kt)("p",null,"Outputs:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-json"},'{\n  "assertionMethod": ["did:snr:123#key-1"],\n  "@context": "https://www.w3.org/ns/did/v1",\n  "controller": "did:snr:123",\n  "id": "did:snr:123",\n  "verificationMethod": [\n    {\n      "controller": "did:snr:123",\n      "id": "did:snr:123#key-1",\n      "publicKeyJwk": {\n        "crv": "P-256",\n        "kty": "EC",\n        "x": "UANQ8pgvJT33JbrnwMiu1L1JCGQFOEm1ThaNAJcFrWA=",\n        "y": "UWm6q5n1iXyeCJLMGDInN40bkkKr8KkoTWDqJBZQXRo="\n      },\n      "type": "JsonWebKey2020"\n    }\n  ]\n}\n')))}h.isMDXComponent=!0}}]);