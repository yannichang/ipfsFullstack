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
<code>npm run build</code>

---
<h2>USAGE</h2>
For store your data, enter the key and data in the form. It should be noted that the key key needs tp be 16, 24 or 32 bytes due to the our encryption design. <br>
A CID will be provided after you click [submit]. <br>
For getting the data you stored, you need the same key and the returned CID. <br>
Your data will be printed!

---

<h2>TEST</h2>
For testing API for the backend (run on localhost:1323), we can also use curl or postman to check the output before building the frontend.<br>
A simple example is shown here. In the terminal, run: <br>
<code>curl -X POST localhost:1323/add -d 'key=abcdabcdabcdabcd' -d 'data=test_message'</code>
