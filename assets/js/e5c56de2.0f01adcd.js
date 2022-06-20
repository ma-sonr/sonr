"use strict";(self.webpackChunksonr_docs=self.webpackChunksonr_docs||[]).push([[704],{3905:function(e,t,n){n.d(t,{Zo:function(){return c},kt:function(){return m}});var r=n(7294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function o(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function s(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},i=Object.keys(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var l=r.createContext({}),d=function(e){var t=r.useContext(l),n=t;return e&&(n="function"==typeof e?e(t):o(o({},t),e)),n},c=function(e){var t=d(e.components);return r.createElement(l.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},u=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,i=e.originalType,l=e.parentName,c=s(e,["components","mdxType","originalType","parentName"]),u=d(n),m=a,h=u["".concat(l,".").concat(m)]||u[m]||p[m]||i;return n?r.createElement(h,o(o({ref:t},c),{},{components:n})):r.createElement(h,o({ref:t},c))}));function m(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var i=n.length,o=new Array(i);o[0]=u;var s={};for(var l in t)hasOwnProperty.call(t,l)&&(s[l]=t[l]);s.originalType=e,s.mdxType="string"==typeof e?e:a,o[1]=s;for(var d=2;d<i;d++)o[d]=n[d];return r.createElement.apply(null,o)}return r.createElement.apply(null,n)}u.displayName="MDXCreateElement"},9449:function(e,t,n){n.r(t),n.d(t,{assets:function(){return c},contentTitle:function(){return l},default:function(){return m},frontMatter:function(){return s},metadata:function(){return d},toc:function(){return p}});var r=n(7462),a=n(3366),i=(n(7294),n(3905)),o=["components"],s={title:"Identifiers",id:"identifiers",displayed_sidebar:"basicsSidebar"},l="Identifiers",d={unversionedId:"guide/advanced/identifiers",id:"guide/advanced/identifiers",title:"Identifiers",description:"Sonr uses the handshake protocol to register subdomains on the .snr/ TLD -- Top-Level-Domain. Using a .snr/ instead of a traditional username was an idea that came to fruition when we realized how much we disliked passwords and not having a genuine identity on the internet. Right now 90% of all social logins on the internet are proxied through Facebook and Google. We believe that ownership of data should reside in the people, not the services that facilitate it.",source:"@site/posts/guide/advanced/identifiers.md",sourceDirName:"guide/advanced",slug:"/guide/advanced/identifiers",permalink:"/posts/guide/advanced/identifiers",editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/posts/guide/advanced/identifiers.md",tags:[],version:"current",frontMatter:{title:"Identifiers",id:"identifiers",displayed_sidebar:"basicsSidebar"},sidebar:"basicsSidebar",previous:{title:"How It Works",permalink:"/posts/guide/how-it-works"},next:{title:"Privacy",permalink:"/posts/guide/advanced/zlY7-privacy"}},c={},p=[{value:"Decentralized Identifiers (DID)",id:"decentralized-identifiers-did",level:3},{value:"Base DID",id:"base-did",level:3},{value:"Controllers",id:"controllers",level:3},{value:"Verification",id:"verification",level:3},{value:"Service DID",id:"service-did",level:3},{value:"Service Endpoints - (Buckets, Channels, Objects, Functions)",id:"service-endpoints---buckets-channels-objects-functions",level:3},{value:"Full Example",id:"full-example",level:3}],u={toc:p};function m(e){var t=e.components,n=(0,a.Z)(e,o);return(0,i.kt)("wrapper",(0,r.Z)({},u,n,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"identifiers"},"Identifiers"),(0,i.kt)("p",null,"Sonr uses the handshake protocol to register subdomains on the .snr/ TLD -- Top-Level-Domain. Using a .snr/ instead of a traditional username was an idea that came to fruition when we realized how much we disliked passwords and not having a genuine identity on the internet. Right now 90% of all social logins on the internet are proxied through Facebook and Google. We believe that ownership of data should reside in the people, not the services that facilitate it."),(0,i.kt)("h3",{id:"decentralized-identifiers-did"},"Decentralized Identifiers (DID)"),(0,i.kt)("p",null,"A DID is a new kind of ID that anyone can use to prove they own their domain. It's similar to a website address, but instead of pointing to a website, it points to you. So if you want to prove that you own a name or an email address, you can get a DID and use it to prove that you own that name or email address."),(0,i.kt)("p",null,"Notable data on each service of the Sonr Protocol is referenced using DID's, making them accessible through your .snr/."),(0,i.kt)("h3",{id:"base-did"},"Base DID"),(0,i.kt)("p",null,"Our base DID (or decentralized identifier) follows a syntactic structure of the ",(0,i.kt)("strong",{parentName:"p"},"root")," (your DID), followed by a ",(0,i.kt)("strong",{parentName:"p"},"method")," (or in our case, in every case with this SDK, Sonr), then followed by a ",(0,i.kt)("strong",{parentName:"p"},"public key")),(0,i.kt)("p",null,(0,i.kt)("img",{parentName:"p",src:"https://archbee-image-uploads.s3.amazonaws.com/YigsjtwFFq_eX7dhChoeN/ze9buUbapxPP7S5ROVXn__6e60b2d-screenshot2022-03-10at25108pm.png",alt:null})),(0,i.kt)("h3",{id:"controllers"},"Controllers"),(0,i.kt)("p",null,(0,i.kt)("strong",{parentName:"p"},"Structure"),"\nControllers are defined by a ",(0,i.kt)("strong",{parentName:"p"},"base DID")," and a fragment"),(0,i.kt)("p",null,(0,i.kt)("strong",{parentName:"p"},"Account Devices"),"\nLinking new devices adds new entries to the users' ",(0,i.kt)("strong",{parentName:"p"},"base DID")),(0,i.kt)("p",null,(0,i.kt)("strong",{parentName:"p"},"IPFS Vault"),"\nUsers' secure IPFS Vault and Mailbox are also listed as controllers on the account's ",(0,i.kt)("strong",{parentName:"p"},"DID Document")),(0,i.kt)("p",null,(0,i.kt)("strong",{parentName:"p"},"Key Management"),"\nIndividual and additional Keys can also be referenced using this fragmented structure."),(0,i.kt)("h3",{id:"verification"},"Verification"),(0,i.kt)("table",null,(0,i.kt)("thead",{parentName:"table"},(0,i.kt)("tr",{parentName:"thead"},(0,i.kt)("th",{parentName:"tr",align:null},(0,i.kt)("strong",{parentName:"th"},"Key")),(0,i.kt)("th",{parentName:"tr",align:null},(0,i.kt)("strong",{parentName:"th"},"Description")))),(0,i.kt)("tbody",{parentName:"table"},(0,i.kt)("tr",{parentName:"tbody"},(0,i.kt)("td",{parentName:"tr",align:null},"Authentication"),(0,i.kt)("td",{parentName:"tr",align:null},"Service used Key.")),(0,i.kt)("tr",{parentName:"tbody"},(0,i.kt)("td",{parentName:"tr",align:null},"Assertion Method"),(0,i.kt)("td",{parentName:"tr",align:null},"Verification key by the user.")),(0,i.kt)("tr",{parentName:"tbody"},(0,i.kt)("td",{parentName:"tr",align:null},"id"),(0,i.kt)("td",{parentName:"tr",align:null},"DID of the verification method")),(0,i.kt)("tr",{parentName:"tbody"},(0,i.kt)("td",{parentName:"tr",align:null},"controller"),(0,i.kt)("td",{parentName:"tr",align:null},"Pointer to the device holding the key.")),(0,i.kt)("tr",{parentName:"tbody"},(0,i.kt)("td",{parentName:"tr",align:null},"type"),(0,i.kt)("td",{parentName:"tr",align:null},"The Verification method type from W3:*   ",(0,i.kt)("inlineCode",{parentName:"td"},"JsonWebkey2020"),(0,i.kt)("inlineCode",{parentName:"td"},"EcdsaSecp256k1RecoveryMethod2020"),(0,i.kt)("inlineCode",{parentName:"td"},"EcdsaSecp256k1VerificationKey2019"))),(0,i.kt)("tr",{parentName:"tbody"},(0,i.kt)("td",{parentName:"tr",align:null},"PublicKey"),(0,i.kt)("td",{parentName:"tr",align:null},"Key format value. The following fields are present:*   ",(0,i.kt)("inlineCode",{parentName:"td"},"publicKeyJwk"))))),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("p",{parentName:"li"},(0,i.kt)("inlineCode",{parentName:"p"},"publicKeyBase58"))),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("p",{parentName:"li"},(0,i.kt)("inlineCode",{parentName:"p"},"publicKeyHex"))),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("p",{parentName:"li"},(0,i.kt)("inlineCode",{parentName:"p"},"blockchainAccountId")," |"))),(0,i.kt)("h3",{id:"service-did"},"Service DID"),(0,i.kt)("p",null,"Apps powered by Sonr will be added to the ",(0,i.kt)("strong",{parentName:"p"},"Blockchain Registry")," with their corresponding ",(0,i.kt)("strong",{parentName:"p"},"Configuration Data"),"."),(0,i.kt)("p",null,(0,i.kt)("img",{parentName:"p",src:"https://archbee-image-uploads.s3.amazonaws.com/YigsjtwFFq_eX7dhChoeN/ZW_uX07qd7Er8odd-Dtkh_1a3639d-screenshot2022-03-10at31119pm.png",alt:null})),(0,i.kt)("h3",{id:"service-endpoints---buckets-channels-objects-functions"},"Service Endpoints - (Buckets, Channels, Objects, Functions)"),(0,i.kt)("p",null,"Modules are resolvable to authorized accounts & services through a subpath to the service extension."),(0,i.kt)("p",null,(0,i.kt)("img",{parentName:"p",src:"https://archbee-image-uploads.s3.amazonaws.com/YigsjtwFFq_eX7dhChoeN/9nerqZTJR7h2HT0y2uVUR_712ad7f-screenshot2022-03-10at31530pm-1.png",alt:null})),(0,i.kt)("h3",{id:"full-example"},"Full Example"),(0,i.kt)("p",null,"Here is a constructed example of a ",(0,i.kt)("inlineCode",{parentName:"p"},"WebAuthn"),"** **token being stored in an accounts ",(0,i.kt)("inlineCode",{parentName:"p"},"DIDDocument")),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-json"},'{\n    "@context": "https://www.w3.org/ns/did/v1",\n    "assertionMethod": ["did:snr:123#key-1"],\n    "alsoKnownAs": ["prad.snr"],\n    "controller": "did:snr:123",\n    "id": "did:snr:123",\n    "verificationMethod": [\n    {\n      "controller": "did:snr:123",\n      "id": "did:snr:123#key-1",\n      "publicKeyJwk": {\n        "crv": "P-256",\n        "kty": "EC",\n        "x": "UANQ8pgvJT33JbrnwMiu1L1JCGQFOEm1ThaNAJcFrWA=",\n        "y": "UWm6q5n1iXyeCJLMGDInN40bkkKr8KkoTWDqJBZQXRo="\n      },\n      "type": "JsonWebKey2020"\n    }\n  ]\n}\n\n')))}m.isMDXComponent=!0}}]);