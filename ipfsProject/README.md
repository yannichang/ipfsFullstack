# IPFS Project
This is a full stack project based on Go and Vue.js. <br>
The objective is to store and retrieve encrypted data from IPFS.
<h2>INSTALL</h2>
<h3>Install Go</h3>
The isntallation of Go can be found in https://golang.org/doc/install.
<h3>go-ipfs-api</h3>
<code>go get -u github.com/ipfs/go-ipfs-api</code><br>
Details about ipfs can be found in https://github.com/ipfs/go-ipfs-api and https://docs.ipfs.io/.
<h3>REST framework </h3>
<code>go get github.com/labstack/echo/v4</code>
More details about this framework can be found in https://github.com/labstack/echo.
<h3>Vue/CLI</h3>
Before installing vue/cli, make sure you installed npm or yarn.<br>
<code>npm install -g @vue/cli</code> <br>
OR <br>
<code>yarn global add @vue/cli</code>
More details about istalling vue/cli and how to create a project can be found in https://cli.vuejs.org/.

---

<h2>SERVER</h2>
<h3>Start IPFS</h3>
<code>ipfs daemon</code>
<h3>Start Backend Server</h3>
<code>go run projectServer.go</code>
<h3>Start Frontend Server</h3>
<code>npm run serve</code> <br>

---
<h2>USAGE</h2>
After the frontend server is running, copy and paste the localhost address in the browser. 
To store your data, please enter the key and data in the form. Note that the key needs to be 16, 24 or 32 bytes due to our encryption design. <br>
A CID will be returned after you click [submit] button. <br>
To get the data you stored from IPFS, you need to provide key and CID. <br>
Your data will be printed!

