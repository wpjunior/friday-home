package api

import (
	"io/ioutil"
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/wpjunior/friday-home/tv"
)

type API interface {
	Run() error
}

type api struct {
	*echo.Echo
	tv tv.TV
}

func New(t tv.TV) API {
	echoInstance := echo.New()
	apiInstance := &api{Echo: echoInstance, tv: t}

	echoInstance.HideBanner = true
	echoInstance.GET("/", apiInstance.index)
	echoInstance.PUT("/api/tv/status", apiInstance.changeTVStatus)

	return apiInstance
}

func (a *api) index(c echo.Context) error {
	return c.HTML(http.StatusOK, `
	<html>
	 <head>
	   <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
	 </head>
	 <body>
	   <h1>Friday home assistant</h1>
	   <button id="turn-on-tv">Turn ON TV</button>
	   <button id="turn-off-tv">Turn OFF TV</button>

	   <script>
		 class TV {
			turnOn() {
			   this.setStatus('on')
			}
			turnOff() {
			   this.setStatus('off')
			}

			async setStatus(status) {
			  const response = await fetch('/api/tv/status', {
				 method: 'PUT',
				 body: status,
			  });
			  const text = await response.text();
			  console.info(text);
			}
		 }
		 const tv = new TV()
		 document.getElementById('turn-on-tv').onclick = () => { tv.turnOn() };
		 document.getElementById('turn-off-tv').onclick = () => { tv.turnOff() };

	   </script>
	 </body>
	</html>
`)
}

func (a *api) changeTVStatus(c echo.Context) error {
	ctx := c.Request().Context()
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	if string(b) == "on" {
		err = a.tv.TurnOn(ctx)
	} else {
		err = a.tv.TurnOff(ctx)
	}
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "OK")
}

func (a *api) Run() error {
	return a.Echo.Start(":8089")
}
