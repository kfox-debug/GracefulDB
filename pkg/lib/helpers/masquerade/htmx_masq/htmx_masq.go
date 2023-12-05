package htmx_masq

var Default string = `
<div class="container-fluid pt-4 px-4">
    <div class="bg-secondary text-center rounded p-4">
		<i class="fa fa-exclamation-triangle"></i> Error: Bad request
    </div>
</div>
`

var Dashboard string = `
<div class="container-fluid pt-4 px-4">
    <div class="bg-secondary text-center rounded p-4">
	    <h4>Dashboard</h4>
    </div>
</div>
`

var Databases string = `
<div class="container-fluid pt-4 px-4">
    <div class="bg-secondary text-center rounded p-4">
	    <h4>Databases</h4>
    </div>
</div>
`

var Accounts string = `
<div class="container-fluid pt-4 px-4">
    <div class="bg-secondary text-center rounded p-4">
	    <h4>Accounts</h4>
    </div>
</div>
`

var Settings string = `
<div class="container-fluid pt-4 px-4">
    <div class="bg-secondary text-center rounded p-4">
	    <h4>Settings</h4>
        <p>The server configuration can be changed via the configuration file.<br>
        Here you can only switch components quickly.</p>
    </div>
</div>

<div class="container-fluid pt-4 px-4">
    <div class="row g-4">
        <div class="col-sm-6 col-xl-3">
            <div class="bg-secondary rounded d-flex align-items-center justify-content-between p-4">
                    <i class="fa fa-info fa-3x text-primary"></i><h5>Basic settings</h5>
                    <div class="ms-3">
                    
                    <p class="mb-2">Environment: <h6 class="mb-0">{{.Env}}</h6></p>
                    <hr>
                    <p class="mb-2">LogPath: <h6 class="mb-0">{{.LogPath}}</h6></p>
                    <hr>
                    <p class="mb-2">ShutdownTimeOut: <h6 class="mb-0">{{.ShutdownTimeOut}}</h6></p>
                    
                    </div>
            </div>
        </div>

        <div class="col-sm-6 col-xl-3">
            <div class="bg-secondary rounded d-flex align-items-center justify-content-between p-4">
                <i class="fa fa-tree fa-3x text-primary"></i><h5>Core settings</h5>
                <div class="ms-3">

                    <p class="mb-2">Bucket size: <h6 class="mb-0">{{.CoreSettings.BucketSize}}</h6></p>

                </div>
            </div>
        </div>
    </div>
</div>

<div class="container-fluid pt-4 px-4">
    <div class="row g-4">
        <div class="col-sm-6 col-xl-3">
            <div class="bg-secondary rounded d-flex align-items-center justify-content-between p-4">
    .WebSocketConnector.Enable: {{.WebSocketConnector.Enable}}<br>
    .WebSocketConnector.Address: {{.WebSocketConnector.Address}}<br>
    .WebSocketConnector.Port: {{.WebSocketConnector.Port}}<br>
    .WebSocketConnector.BufferSize.Read: {{.WebSocketConnector.BufferSize.Read}}<br>
    .WebSocketConnector.BufferSize.Write: {{.WebSocketConnector.BufferSize.Write}}<br> <br>
            </div>
        </div>
    
        <div class="col-sm-6 col-xl-3">
        <div class="bg-secondary rounded d-flex align-items-center justify-content-between p-4">
    .RestConnector.Enable: {{.RestConnector.Enable}}<br>
    .RestConnector.Address: {{.RestConnector.Address}}<br>
    .RestConnector.Port: {{.RestConnector.Port}}<br> <br>
            </div>
        </div>

        <div class="col-sm-6 col-xl-3">
            <div class="bg-secondary rounded d-flex align-items-center justify-content-between p-4">
    .GrpcConnector.Enable: {{.GrpcConnector.Enable}}<br>
    .GrpcConnector.Address: {{.GrpcConnector.Address}}<br>
    .GrpcConnector.Port: {{.GrpcConnector.Port}}<br> <br>
            </div>
        </div>
    
        <div class="col-sm-6 col-xl-3">
            <div class="bg-secondary rounded d-flex align-items-center justify-content-between p-4">
    .WebServer.Enable: {{.WebServer.Enable}}<br>
    .WebServer.Address: {{.WebServer.Address}}<br>
    .WebServer.Port: {{.WebServer.Port}}<br> <br>
            </div>
        </div>
    
    </div>
</div>
`
