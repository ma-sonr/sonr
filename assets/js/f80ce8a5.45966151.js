"use strict";(self.webpackChunksonr_docs=self.webpackChunksonr_docs||[]).push([[3655],{3905:function(e,t,r){r.d(t,{Zo:function(){return u},kt:function(){return y}});var n=r(7294);function a(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function i(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function o(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?i(Object(r),!0).forEach((function(t){a(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):i(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function c(e,t){if(null==e)return{};var r,n,a=function(e,t){if(null==e)return{};var r,n,a={},i=Object.keys(e);for(n=0;n<i.length;n++)r=i[n],t.indexOf(r)>=0||(a[r]=e[r]);return a}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(n=0;n<i.length;n++)r=i[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(a[r]=e[r])}return a}var d=n.createContext({}),s=function(e){var t=n.useContext(d),r=t;return e&&(r="function"==typeof e?e(t):o(o({},t),e)),r},u=function(e){var t=s(e.components);return n.createElement(d.Provider,{value:t},e.children)},l={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},p=n.forwardRef((function(e,t){var r=e.components,a=e.mdxType,i=e.originalType,d=e.parentName,u=c(e,["components","mdxType","originalType","parentName"]),p=s(r),y=a,v=p["".concat(d,".").concat(y)]||p[y]||l[y]||i;return r?n.createElement(v,o(o({ref:t},u),{},{components:r})):n.createElement(v,o({ref:t},u))}));function y(e,t){var r=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var i=r.length,o=new Array(i);o[0]=p;var c={};for(var d in t)hasOwnProperty.call(t,d)&&(c[d]=t[d]);c.originalType=e,c.mdxType="string"==typeof e?e:a,o[1]=c;for(var s=2;s<i;s++)o[s]=r[s];return n.createElement.apply(null,o)}return n.createElement.apply(null,r)}p.displayName="MDXCreateElement"},9800:function(e,t,r){r.r(t),r.d(t,{assets:function(){return u},contentTitle:function(){return d},default:function(){return y},frontMatter:function(){return c},metadata:function(){return s},toc:function(){return l}});var n=r(7462),a=r(3366),i=(r(7294),r(3905)),o=["components"],c={title:"Privacy",slug:"zlY7-privacy",createdAt:new Date("2022-04-15T20:04:42.000Z"),updatedAt:new Date("2022-04-23T19:51:52.000Z")},d="Privacy",s={unversionedId:"guide/advanced/privacy",id:"guide/advanced/privacy",title:"Privacy",description:"Overview",source:"@site/posts/guide/advanced/privacy.md",sourceDirName:"guide/advanced",slug:"/guide/advanced/zlY7-privacy",permalink:"/posts/guide/advanced/zlY7-privacy",editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/posts/guide/advanced/privacy.md",tags:[],version:"current",frontMatter:{title:"Privacy",slug:"zlY7-privacy",createdAt:"2022-04-15T20:04:42.000Z",updatedAt:"2022-04-23T19:51:52.000Z"},sidebar:"basicsSidebar",previous:{title:"Identifiers",permalink:"/posts/guide/advanced/identifiers"},next:{title:"Security",permalink:"/posts/guide/advanced/security"}},u={},l=[{value:"Overview",id:"overview",level:2},{value:"Encrypted Messaging",id:"encrypted-messaging",level:3},{value:"Decentralized Data Storage",id:"decentralized-data-storage",level:3}],p={toc:l};function y(e){var t=e.components,r=(0,a.Z)(e,o);return(0,i.kt)("wrapper",(0,n.Z)({},p,r,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"privacy"},"Privacy"),(0,i.kt)("h2",{id:"overview"},"Overview"),(0,i.kt)("p",null,"Our protocols are built on top of block chain technology, letting us take advantage of it's distributive anonymous nature. All data stored is denormalized. Writes on our chain are only related by did, and publicKey. With each message sent on our network handshaked, with TCP/TLS."," "),(0,i.kt)("h3",{id:"encrypted-messaging"},"Encrypted Messaging"),(0,i.kt)("p",null,"All data stored, communicated on our network is end to end encrypted. Using ",(0,i.kt)("inlineCode",{parentName:"p"},"EDC25519")," elliptic curve standard, each operation requires a key exchange inorder to operate on user data. This makes it so your data can only opened by you or users you trust."),(0,i.kt)("h3",{id:"decentralized-data-storage"},"Decentralized Data Storage"),(0,i.kt)("p",null,"We use ",(0,i.kt)("inlineCode",{parentName:"p"},"IPFS")," to store ",(0,i.kt)("inlineCode",{parentName:"p"},"Blobs"),", and ",(0,i.kt)("inlineCode",{parentName:"p"},"Objects"),". Each item stored is sharded, meaning your data is stored in fragments across our network. Each item is encrypted using the ",(0,i.kt)("inlineCode",{parentName:"p"},"EDC25519")," standard when stored on the nodes file system. Each ",(0,i.kt)("inlineCode",{parentName:"p"},"Highway")," node is associated with an IPFS instance meaning your data can be sharded across our network with high Encryption standards."),(0,i.kt)("h1",{id:""}))}y.isMDXComponent=!0}}]);