package api

import (
	"io/ioutil"
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/wpjunior/friday-home/player"
	"github.com/wpjunior/friday-home/tv"
)

type API interface {
	Run() error
}

type api struct {
	*echo.Echo
	tv     tv.TV
	player player.Player
}

func New(t tv.TV, p player.Player) API {
	echoInstance := echo.New()
	apiInstance := &api{Echo: echoInstance, tv: t, player: p}

	echoInstance.HideBanner = true
	echoInstance.GET("/", apiInstance.index)
	echoInstance.GET("/api/tv/profiles", apiInstance.getProfiles)
	echoInstance.GET("/api/tv/volume/up", apiInstance.volumeUp)
	echoInstance.GET("/api/tv/volume/down", apiInstance.volumeDown)
	echoInstance.PUT("/api/tv/status", apiInstance.changeTVStatus)

	return apiInstance
}

func (a *api) index(c echo.Context) error {
	return c.HTML(http.StatusOK, `
	<html>
	 <head>
	   <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
	   <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
	   <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/css/materialize.min.css">
	 </head>
	 <body>
	   <nav>
		 <div class="nav-wrapper">
		   <a href="#" class="brand-logo">Friday</a>
		 </div>
	   </nav>

	   <pre id="console"></pre>
	   <button id="volume-up">Volume up</button>
	   <button id="volume-down">Volume down</button>
	   <ul id="profiles" class="collection">
	   </ul>

	   <script>
		 const consoleDOM = document.getElementById('console');
		 async function setStatus(status) {
			 consoleDOM.textContent = "Loading ..."
			 const response = await fetch('/api/tv/status', {
				 method: 'PUT',
				 body: status,
			  });
			  if (response.status === 200) {
				 consoleDOM.textContent = ""
			  } else {
				 const text = await response.text();
				 consoleDOM.textContent = text;
			  }
		 }

		 document.getElementById('volume-up').onclick = async () => {
			  await fetch('/api/tv/volume/up')
		 }

		 document.getElementById('volume-down').onclick = async () => {
			  await fetch('/api/tv/volume/down')
		 }

		 async function loadProfiles() {
			const response = await fetch('/api/tv/profiles');
			const profiles = await response.json()

			const profilesDOM = document.getElementById('profiles');
			for (const profile of profiles) {
			   const li = document.createElement('li');
			   li.className = "collection-item avatar";
			   li.onclick = setStatus.bind(null, profile.id);

			   const i = document.createElement('i');
			   i.className = "material-icons circle green";
			   i.textContent = "insert_chart";
			   li.appendChild(i);

			   const span = document.createElement('span')
			   span.className = "title";
			   span.textContent = profile.name
			   li.appendChild(span)

			   profilesDOM.appendChild(li);
			}
		 }

		 loadProfiles()
	   </script>
	 </body>
	</html>
`)
}

func (a *api) volumeUp(c echo.Context) error {
	return a.tv.VolumeUp(c.Request().Context())
}

func (a *api) volumeDown(c echo.Context) error {
	return a.tv.VolumeDown(c.Request().Context())
}

func (a *api) changeTVStatus(c echo.Context) error {
	ctx := c.Request().Context()
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	if string(b) == "off" {
		err = a.tv.TurnOff(ctx)
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, "OK")
	}

	profile := GetProfile(string(b))
	if profile == nil {
	} else {
		err = a.tv.TurnOn(ctx)
		if err == nil {
			err = a.player.PlayYoutubeChannel(profile.YoutubeChannels[0])
		}
	}
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "OK")
}

func (a *api) getProfiles(c echo.Context) error {
	return c.JSON(http.StatusOK, Profiles)
}

func (a *api) Run() error {
	return a.Echo.Start(":8089")
}
