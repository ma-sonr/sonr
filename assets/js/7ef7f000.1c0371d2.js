"use strict";(self.webpackChunksonr_docs=self.webpackChunksonr_docs||[]).push([[2475],{3905:function(e,t,n){n.d(t,{Zo:function(){return d},kt:function(){return h}});var o=n(7294);function r(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);t&&(o=o.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,o)}return n}function a(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){r(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function s(e,t){if(null==e)return{};var n,o,r=function(e,t){if(null==e)return{};var n,o,r={},i=Object.keys(e);for(o=0;o<i.length;o++)n=i[o],t.indexOf(n)>=0||(r[n]=e[n]);return r}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(o=0;o<i.length;o++)n=i[o],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(r[n]=e[n])}return r}var c=o.createContext({}),l=function(e){var t=o.useContext(c),n=t;return e&&(n="function"==typeof e?e(t):a(a({},t),e)),n},d=function(e){var t=l(e.components);return o.createElement(c.Provider,{value:t},e.children)},u={inlineCode:"code",wrapper:function(e){var t=e.children;return o.createElement(o.Fragment,{},t)}},p=o.forwardRef((function(e,t){var n=e.components,r=e.mdxType,i=e.originalType,c=e.parentName,d=s(e,["components","mdxType","originalType","parentName"]),p=l(n),h=r,f=p["".concat(c,".").concat(h)]||p[h]||u[h]||i;return n?o.createElement(f,a(a({ref:t},d),{},{components:n})):o.createElement(f,a({ref:t},d))}));function h(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var i=n.length,a=new Array(i);a[0]=p;var s={};for(var c in t)hasOwnProperty.call(t,c)&&(s[c]=t[c]);s.originalType=e,s.mdxType="string"==typeof e?e:r,a[1]=s;for(var l=2;l<i;l++)a[l]=n[l];return o.createElement.apply(null,a)}return o.createElement.apply(null,n)}p.displayName="MDXCreateElement"},7784:function(e,t,n){n.r(t),n.d(t,{assets:function(){return d},contentTitle:function(){return c},default:function(){return h},frontMatter:function(){return s},metadata:function(){return l},toc:function(){return u}});var o=n(7462),r=n(3366),i=(n(7294),n(3905)),a=["components"],s={title:"Discovery",id:"discovery",displayed_sidebar:"modulesSidebar"},c="Discovery",l={unversionedId:"build-apps/motor/discovery",id:"build-apps/motor/discovery",title:"Discovery",description:"The Sonr protocol uses various methods and fallbacks to ensure no friction discoverability when finding another user. Its node has an Exchange service that provides immediate discovery, validation, and verification access for every peer. The service operates in two separate modes: Local and Global, with the Local mode functioning similar to Airdrop and Global to Email.",source:"@site/posts/build-apps/motor/discovery.md",sourceDirName:"build-apps/motor",slug:"/build-apps/motor/discovery",permalink:"/posts/build-apps/motor/discovery",editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/posts/build-apps/motor/discovery.md",tags:[],version:"current",frontMatter:{title:"Discovery",id:"discovery",displayed_sidebar:"modulesSidebar"},sidebar:"modulesSidebar",previous:{title:"Access & Authentication",permalink:"/posts/build-apps/motor/access-authentication"},next:{title:"Transmission",permalink:"/posts/build-apps/motor/transmission"}},d={},u=[],p={toc:u};function h(e){var t=e.components,n=(0,r.Z)(e,a);return(0,i.kt)("wrapper",(0,o.Z)({},p,n,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"discovery"},"Discovery"),(0,i.kt)("p",null,"The Sonr protocol uses various methods and fallbacks to ensure no friction discoverability when finding another user. Its node has an Exchange service that provides immediate discovery, validation, and verification access for every peer. The service operates in two separate modes: Local and Global, with the Local mode functioning similar to Airdrop and Global to Email."),(0,i.kt)("h1",{id:"local"},"Local"),(0,i.kt)("p",null,"Before the discovery process begins, nodes provide their Lat/Lng to join the local lobby - when location isn't provided, the node will operate on Global only mode. After receiving the values, an OLC (Open Location Code) is generated, which becomes the suffix for the new Local Topic. Peers that are validated then have access to Read/Write on the Local Room's DHT (Distributed Hash Table), where they would provide any Profile Info or Device Metadata."),(0,i.kt)("p",null,"Facilitating this is a combination of transports including mDNS, Bluetooth, Rendevouz Signalling, and the DHT.\nAfter Peers are discovered, the cached copy of the DHT is streamed to any client-facing implementation bound to the node. The peers are then displayed on the UI."),(0,i.kt)("h1",{id:"global"},"Global"),(0,i.kt)("p",null,"The Global topic uses a lighter data structure for managing peers than rather the robust DHT used locally. To solve the issue of scale, we created Beam - a decentralized simple key/value store that has verifiable write access. By using Beam in combination with the node's HDNS Subdomain, we can successfully query the entire network in an instant, even when the node itself is offline."),(0,i.kt)("h1",{id:"offline"},"Offline"),(0,i.kt)("p",null,"A common issue with P2P systems is handling the situation when the node is offline. With Sonr, we solve this issue by implementing the DIDCommMessaging Spec along with an IPFS mailbox where only users can access with their .snr/."))}h.isMDXComponent=!0}}]);