package socket_io

import (
	"github.com/Baiguoshuai1/shadiaosocketio"
	"github.com/Baiguoshuai1/shadiaosocketio/websocket"
	patient "github.com/cuida-me/mvp-backend/internal/application/patient/contracts"
	dto "github.com/cuida-me/mvp-backend/internal/application/patient/dto"
	"log"
)

type websocketConnection struct {
	newPatientSessionUseCase patient.NewPatientSession
	refreshSessionQRUseCase  patient.RefreshSessionQR
}

func NewWebsocketConnection(newPatientSessionUseCase patient.NewPatientSession, refreshSessionQRUseCase patient.RefreshSessionQR) *websocketConnection {
	return &websocketConnection{
		newPatientSessionUseCase: newPatientSessionUseCase,
		refreshSessionQRUseCase:  refreshSessionQRUseCase,
	}
}

func (ws websocketConnection) SocketConnection() *shadiaosocketio.Server {
	server := shadiaosocketio.NewServer(*websocket.GetDefaultWebsocketTransport())

	server.On(shadiaosocketio.OnConnection, func(c *shadiaosocketio.Channel) {
		log.Println("[server] connected! id:", c.Id())
	})

	server.On(shadiaosocketio.OnDisconnection, func(c *shadiaosocketio.Channel, reason websocket.CloseError) {
		log.Println("[server] received disconnect", c.Id(), "code:", reason.Code, "text:", reason.Text)
	})

	server.On(shadiaosocketio.OnError, func(c *shadiaosocketio.Channel, arg1 string) {
		log.Println("[server] received message:", "arg1:", arg1)
	})

	server.On("patient-qr", ws.newPatientSession)

	server.On("patient-qr-refresh", ws.refreshSessionQR)

	server.On("schedules-screen", ws.schedulesScreen)

	server.On(shadiaosocketio.OnMessage, func(c *shadiaosocketio.Channel, arg1 string) {
		log.Println("[server] received message:", "arg1:", arg1)
	})

	return server
}

func (ws websocketConnection) newPatientSession(c *shadiaosocketio.Channel, request dto.NewPatientSessionRequest) {
	response, err := ws.newPatientSessionUseCase.Execute(c.Request().Context(), &request, c.Id())
	if err != nil {
		log.Println(err)
	}
	c.Emit("patient-qr", response)
}

func (ws websocketConnection) refreshSessionQR(c *shadiaosocketio.Channel, request dto.RefreshSessionQRRequest) {
	response, err := ws.refreshSessionQRUseCase.Execute(c.Request().Context(), &request, c.Id())
	if err != nil {
		log.Println(err)
	}
	c.Emit("patient-qr", response)
}

func (ws websocketConnection) schedulesScreen(c *shadiaosocketio.Channel) {
	c.Emit("schedules-screen")
}
