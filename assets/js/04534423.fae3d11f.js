"use strict";(self.webpackChunksonr_docs=self.webpackChunksonr_docs||[]).push([[3072],{3905:function(e,t,n){n.d(t,{Zo:function(){return d},kt:function(){return h}});var r=n(7294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function o(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function l(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},i=Object.keys(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var s=r.createContext({}),c=function(e){var t=r.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):o(o({},t),e)),n},d=function(e){var t=c(e.components);return r.createElement(s.Provider,{value:t},e.children)},u={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},p=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,i=e.originalType,s=e.parentName,d=l(e,["components","mdxType","originalType","parentName"]),p=c(n),h=a,f=p["".concat(s,".").concat(h)]||p[h]||u[h]||i;return n?r.createElement(f,o(o({ref:t},d),{},{components:n})):r.createElement(f,o({ref:t},d))}));function h(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var i=n.length,o=new Array(i);o[0]=p;var l={};for(var s in t)hasOwnProperty.call(t,s)&&(l[s]=t[s]);l.originalType=e,l.mdxType="string"==typeof e?e:a,o[1]=l;for(var c=2;c<i;c++)o[c]=n[c];return r.createElement.apply(null,o)}return r.createElement.apply(null,n)}p.displayName="MDXCreateElement"},6938:function(e,t,n){n.r(t),n.d(t,{assets:function(){return d},contentTitle:function(){return s},default:function(){return h},frontMatter:function(){return l},metadata:function(){return c},toc:function(){return u}});var r=n(7462),a=n(3366),i=(n(7294),n(3905)),o=["components"],l={title:"Interface",id:"interface",displayed_sidebar:"modulesSidebar"},s="Motor Interface",c={unversionedId:"build-apps/motor/interface",id:"build-apps/motor/interface",title:"Interface",description:"Overview",source:"@site/posts/build-apps/motor/interface.md",sourceDirName:"build-apps/motor",slug:"/build-apps/motor/interface",permalink:"/posts/build-apps/motor/interface",editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/posts/build-apps/motor/interface.md",tags:[],version:"current",frontMatter:{title:"Interface",id:"interface",displayed_sidebar:"modulesSidebar"},sidebar:"modulesSidebar",previous:{title:"Getting Started",permalink:"/posts/build-apps/installation"},next:{title:"Access & Authentication",permalink:"/posts/build-apps/motor/access-authentication"}},d={},u=[{value:"Overview",id:"overview",level:2},{value:"Interaction Methods",id:"interaction-methods",level:3},{value:"Highway Methods",id:"highway-methods",level:3},{value:"Node Methods",id:"node-methods",level:3},{value:"AccessApp",id:"accessapp",level:3},{value:"InteractObject",id:"interactobject",level:3},{value:"ListenChannel",id:"listenchannel",level:3},{value:"LinkDevice",id:"linkdevice",level:3},{value:"ReadBucket",id:"readbucket",level:3},{value:"ReadObject",id:"readobject",level:3},{value:"RequestBucket",id:"requestbucket",level:3},{value:"RespondBucket",id:"respondbucket",level:3},{value:"ClearMessage()",id:"clearmessage",level:3},{value:"DownloadVault",id:"downloadvault",level:3},{value:"OpenMailbox",id:"openmailbox",level:3},{value:"RequestPeer",id:"requestpeer",level:3},{value:"RespondPeer",id:"respondpeer",level:3},{value:"SendMessage",id:"sendmessage",level:3},{value:"UploadVault",id:"uploadvault",level:3}],p={toc:u};function h(e){var t=e.components,n=(0,a.Z)(e,o);return(0,i.kt)("wrapper",(0,r.Z)({},p,n,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"motor-interface"},"Motor Interface"),(0,i.kt)("h2",{id:"overview"},"Overview"),(0,i.kt)("h3",{id:"interaction-methods"},"Interaction Methods"),(0,i.kt)("p",null,"These are methods that are handled only between the user and their data, or directly with another peer."),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-Text"},"\n ClearMessage()\n DownloadVault()\n OpenMailbox()\n RequestPeer()\n RespondPeer()\n SendMessage()\n UploadVault()\n")),(0,i.kt)("h3",{id:"highway-methods"},"Highway Methods"),(0,i.kt)("p",null,"These methods interact with a Sonr Highway node which communicates with the underlying blockchain"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-Text"}," AccessApp()\n InteractObject()\n ListenChannel()\n LinkDevice()\n LinkPeer()\n ReadBucket()\n ReadObject()\n RequestBucket()\n RespondBucket()\n")),(0,i.kt)("h3",{id:"node-methods"},"Node Methods"),(0,i.kt)("p",null,"These are developer facing methods that manage the state of the underlying motor node."),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-Text"},"\n Start()\n Stop()\n Pause()\n Resume()\n Status()\n")),(0,i.kt)("h3",{id:"accessapp"},"AccessApp"),(0,i.kt)("p",null,"User authenticates a Registered Application on Sonr with their DID Based Multisignature key for all their devices. Creates a new Bucket inside the User bucket for the newly provisioned Application."),(0,i.kt)("h3",{id:"interactobject"},"InteractObject"),(0,i.kt)("p",null,"Users map the new data for a specific type definition presented in the UI, and push the updated data to the corresponding application in their Bucket. This utilizes the JWE process in order to encrypt data from the User end."),(0,i.kt)("h3",{id:"listenchannel"},"ListenChannel"),(0,i.kt)("p",null,"User specifies which application data stream to begin reading for data. The returned channel is a listenable stream or callback depending on Device architecture."),(0,i.kt)("h3",{id:"linkdevice"},"LinkDevice"),(0,i.kt)("p",null,"This method allows motor based applications to link an addition WebAuthN credential to their top-level"),(0,i.kt)("h3",{id:"readbucket"},"ReadBucket"),(0,i.kt)("p",null,"This method be utilized with pulling a Application Specific buckets type definitions, functions, etc. In order to render the payloads onto the frontend UI."),(0,i.kt)("h3",{id:"readobject"},"ReadObject"),(0,i.kt)("p",null,"This method begins the request process for reading the individual object values of another User's app bucket. Provisioned users will automatically get access (if they called RequestBucket already) and unprovisioned users will send the RequestBucket to the corresponding peer."),(0,i.kt)("h3",{id:"requestbucket"},"RequestBucket"),(0,i.kt)("p",null,"This method is utilized for accessing another users application specific data holistically. This would be utilized for full access to a peers App Data config"),(0,i.kt)("h3",{id:"respondbucket"},"RespondBucket"),(0,i.kt)("p",null,"This method is utilized for responding to a request from another user from the mailbox folder in their User specific bucket."),(0,i.kt)("h3",{id:"clearmessage"},"ClearMessage()"),(0,i.kt)("p",null,"Removes a message from the mailbox with provided ID."),(0,i.kt)("h3",{id:"downloadvault"},"DownloadVault"),(0,i.kt)("p",null,"Retreives a file or buffer from the user data vault from in their Bucket - with either a specified CID or REGEX Match string for querying files."),(0,i.kt)("h3",{id:"openmailbox"},"OpenMailbox"),(0,i.kt)("p",null,"Retreives the list of messages stored in the user Mailbox, with the id and their corresponding buffer body."),(0,i.kt)("h3",{id:"requestpeer"},"RequestPeer"),(0,i.kt)("p",null,"Adds a connected peer into the root level DID as a InvocationMethod in order to have ongoing connection."),(0,i.kt)("h3",{id:"respondpeer"},"RespondPeer"),(0,i.kt)("p",null,"Approves/Declines a request from another peer to connect as an InvocationMethod on the top level DID Document for a User."),(0,i.kt)("h3",{id:"sendmessage"},"SendMessage"),(0,i.kt)("p",null,"Send a message to another peer's mailbox utilizing any arbitrary buffer of data, signed with the users DIDDocument. The other user then fetches the public key from the sending user's DIDDocument on the sonr blockchain in order to verify the authenticity of the sender."),(0,i.kt)("h3",{id:"uploadvault"},"UploadVault"),(0,i.kt)("p",null,"Uploads data to the User specific bucket in their Vault directory."))}h.isMDXComponent=!0}}]);