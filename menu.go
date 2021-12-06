package main

import (
	"adarocket/rocket/client"
	"fmt"
	"github.com/adarocket/proto/proto-gen/cardano"
	"google.golang.org/grpc"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var informClient *client.ControllerClient
var authClient *client.AuthClient
var cardanoClient *client.CardanoClient

// MenuField -
type MenuField struct {
	Title, UUID, Status string
	View                func(w fyne.Window, uuid string) fyne.CanvasObject
}

// MenuFields -
var MenuFields = map[string]*MenuField{
	"welcome": {
		Title:  "Welcome",
		UUID:   "",
		Status: "",
		View:   welcomeScreen,
	},
}

// MenuIndex -
var MenuIndex = map[string][]string{
	"": {"welcome"},
}

func welcomeScreen(w fyne.Window, uuid string) fyne.CanvasObject {
	return container.NewVBox()
}

func authMethods() map[string]bool {
	return map[string]bool{
		"/cardano.Cardano/" + "GetStatistic":  true, //cardano.Cardano
		"/Common.Controller/" + "GetNodeList": true, //Common.Controller
	}
}

func setupInterceptorAndClient(accessToken, serverURL string) {
	transportOption := grpc.WithInsecure()

	interceptor, err := client.NewAuthInterceptor(authMethods(), accessToken)
	if err != nil {
		log.Fatal("cannot create auth interceptor: ", err)
	}

	clientConn, err := grpc.Dial(serverURL, transportOption, grpc.WithUnaryInterceptor(interceptor.Unary()))
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	informClient = client.NewControllerClient(clientConn)
	cardanoClient = client.NewCardanoClient(clientConn)
}

func informationScreen(w fyne.Window, uuid string) fyne.CanvasObject {
	intro := widget.NewLabel("No information about this node")
	intro.Wrapping = fyne.TextWrapWord

	resp := nodeInfoMap[uuid]

	var items []fyne.CanvasObject
	item := new(widget.Form)

	/*if resp.NodeAuthData != nil {
		item = widget.NewForm(
			widget.NewFormItem("Node Auth Data", widget.NewLabel("")),

			widget.NewFormItem("Ticker", widget.NewLabel(resp.NodeAuthData.Ticker)),
			widget.NewFormItem("Uuid", widget.NewLabel(resp.NodeAuthData.Uuid)),
			widget.NewFormItem("Blockchain", widget.NewLabel(resp.NodeAuthData.Blockchain)),
		)

		items = append(items, item, widget.NewSeparator())
	}*/

	if resp.Statistic.NodeBasicData != nil {
		item = widget.NewForm(
			widget.NewFormItem("Node Basic Data", widget.NewLabel("")),

			widget.NewFormItem("Ticker", widget.NewLabel(resp.Statistic.NodeBasicData.Ticker)),
			widget.NewFormItem("Type", widget.NewLabel(resp.Statistic.NodeBasicData.Type)),
			widget.NewFormItem("Location", widget.NewLabel(resp.Statistic.NodeBasicData.Location)),
			widget.NewFormItem("Node version", widget.NewLabel(resp.Statistic.NodeBasicData.NodeVersion)),
		)

		items = append(items, item, widget.NewSeparator())
	}

	if resp.Statistic.ServerBasicData != nil {
		item = widget.NewForm(
			widget.NewFormItem("Server Basic Data", widget.NewLabel("")),

			widget.NewFormItem("IPv4", widget.NewLabel(resp.Statistic.ServerBasicData.Ipv4)),
			widget.NewFormItem("IPv6", widget.NewLabel(resp.Statistic.ServerBasicData.Ipv6)),
			widget.NewFormItem("Linux name", widget.NewLabel(resp.Statistic.ServerBasicData.LinuxName)),
			widget.NewFormItem("Linux version", widget.NewLabel(resp.Statistic.ServerBasicData.LinuxVersion)),
		)
		items = append(items, item, widget.NewSeparator())
	}

	if resp.Statistic.Online != nil {
		item = widget.NewForm(
			widget.NewFormItem("Online", widget.NewLabel("")),

			widget.NewFormItem("Since start", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.Online.SinceStart))),
			widget.NewFormItem("Pings", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.Online.Pings))),
			widget.NewFormItem("Node active", widget.NewLabel(fmt.Sprintf("%t", resp.Statistic.Online.NodeActive))),
			widget.NewFormItem("Node active pings", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.Online.NodeActivePings))),
		)
		items = append(items, item, widget.NewSeparator())
	}

	if resp.Statistic.MemoryState != nil {
		item = widget.NewForm(
			widget.NewFormItem("Memory", widget.NewLabel("")),

			widget.NewFormItem("Total", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.MemoryState.Total))),
			widget.NewFormItem("Used", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.MemoryState.Used))),
			widget.NewFormItem("Buffers", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.MemoryState.Buffers))),
			widget.NewFormItem("Cached", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.MemoryState.Cached))),
			widget.NewFormItem("Free", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.MemoryState.Free))),
			widget.NewFormItem("Available", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.MemoryState.Available))),
			widget.NewFormItem("Active", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.MemoryState.Active))),
			widget.NewFormItem("Inactive", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.MemoryState.Inactive))),
			widget.NewFormItem("Swap Total", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.MemoryState.SwapTotal))),
			widget.NewFormItem("Swap Used", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.MemoryState.SwapUsed))),
			widget.NewFormItem("Swap Cached", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.MemoryState.SwapCached))),
			widget.NewFormItem("Swap Free", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.MemoryState.SwapFree))),
			widget.NewFormItem("Mem Available Enabled", widget.NewLabel(fmt.Sprintf("%t", resp.Statistic.MemoryState.MemAvailableEnabled))),
		)
		items = append(items, item, widget.NewSeparator())
	}

	if resp.Statistic.CpuState != nil {
		item = widget.NewForm(
			widget.NewFormItem("CPU state", widget.NewLabel("")),

			widget.NewFormItem("CPU Qty", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.CpuState.CpuQty))),
			widget.NewFormItem("Average workload", widget.NewLabel(fmt.Sprintf("%f", resp.Statistic.CpuState.AverageWorkload))),
		)
		items = append(items, item, widget.NewSeparator())
	}

	if resp.Statistic.Epoch != nil {
		item = widget.NewForm(
			widget.NewFormItem("Epoch", widget.NewLabel("")),

			widget.NewFormItem("Epoch number", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.Epoch.EpochNumber))),
		)
		items = append(items, item, widget.NewSeparator())
	}

	if resp.Statistic.NodeState != nil {
		item = widget.NewForm(
			widget.NewFormItem("Node State", widget.NewLabel("")),
			widget.NewFormItem("Tip diff", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.NodeState.TipDiff))),
			widget.NewFormItem("Density", widget.NewLabel(fmt.Sprintf("%f", resp.Statistic.NodeState.Density))),
		)
		items = append(items, item, widget.NewSeparator())
	}

	if resp.Statistic.NodePerformance != nil {
		item = widget.NewForm(
			widget.NewFormItem("Node Performance", widget.NewLabel("")),

			widget.NewFormItem("Processed Tx", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.NodePerformance.ProcessedTx))),
			widget.NewFormItem("Peers In", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.NodePerformance.PeersIn))),
			widget.NewFormItem("Peers Out", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.NodePerformance.PeersOut))),
		)
		items = append(items, item, widget.NewSeparator())
	}

	if resp.Statistic.KesData != nil {
		item = widget.NewForm(
			widget.NewFormItem("KES Data", widget.NewLabel("")),

			widget.NewFormItem("KES current", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.KesData.KesCurrent))),
			widget.NewFormItem("KES remaining", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.KesData.KesRemaining))),
			widget.NewFormItem("KES exp date", widget.NewLabel(resp.Statistic.KesData.KesExpDate)),
		)
		items = append(items, item, widget.NewSeparator())
	}

	if resp.Statistic.Blocks != nil {
		item = widget.NewForm(
			widget.NewFormItem("Blocks", widget.NewLabel("")),

			widget.NewFormItem("Block leader", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.Blocks.BlockLeader))),
			widget.NewFormItem("Block adopted", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.Blocks.BlockAdopted))),
			widget.NewFormItem("Block invalid", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.Blocks.BlockInvalid))),
		)
		items = append(items, item, widget.NewSeparator())
	}

	if resp.Statistic.Updates != nil {
		item = widget.NewForm(
			widget.NewFormItem("Updates", widget.NewLabel("")),

			widget.NewFormItem("Informer actual", widget.NewLabel(resp.Statistic.Updates.InformerActual)),
			widget.NewFormItem("Informer available", widget.NewLabel(resp.Statistic.Updates.InformerAvailable)),
			widget.NewFormItem("Updater actual", widget.NewLabel(resp.Statistic.Updates.UpdaterActual)),
			widget.NewFormItem("Updater available", widget.NewLabel(resp.Statistic.Updates.UpdaterAvailable)),
			widget.NewFormItem("Packages available", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.Updates.PackagesAvailable))),
		)
		items = append(items, item, widget.NewSeparator())
	}

	if resp.Statistic.Security != nil {
		item = widget.NewForm(
			widget.NewFormItem("Security", widget.NewLabel("")),

			widget.NewFormItem("SSH Attack Attempts", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.Security.SshAttackAttempts))),
			widget.NewFormItem("Security Packages Available", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.Security.SecurityPackagesAvailable))),

			widget.NewFormItem("Security", widget.NewLabel("")),
		)
		items = append(items, item, widget.NewSeparator())
	}

	if resp.Statistic.StakeInfo != nil {
		item = widget.NewForm(
			widget.NewFormItem("StakeInfo", widget.NewLabel("")),

			widget.NewFormItem("Live stake", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.StakeInfo.LiveStake))),
			widget.NewFormItem("Active stake", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.StakeInfo.ActiveStake))),
			widget.NewFormItem("Pledge", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.StakeInfo.Pledge))),
		)
		items = append(items, item, widget.NewSeparator())
	}

	if len(items) > 0 {
		// return container.NewVScroll(container.NewVBox(widget.NewAccordion(items...)))
		// return container.NewVScroll(container.NewVBox(items...))

		return container.NewVScroll(container.NewVBox(items...))
	}

	// fyne.LogError("Items les then 0", errors.New("Items les then 0"))
	return container.NewVBox(intro)
}

func loginScreen(w fyne.Window, a fyne.App) fyne.CanvasObject {
	a.Settings().SetTheme(theme.DarkTheme())

	usernameField := widget.NewEntry()
	passwordField := widget.NewPasswordEntry()

	return container.NewCenter(
		container.NewVBox(
			widget.NewForm(
				widget.NewFormItem("User name", usernameField),
				widget.NewFormItem("Password", passwordField),
			),
			container.NewHBox(
				widget.NewButtonWithIcon("Cancel", theme.CancelIcon(), func() {
					a.Quit()
				}),

				widget.NewButtonWithIcon("Submit", theme.ConfirmIcon(), func() {
					token, err := authClient.Login(usernameField.Text, passwordField.Text)
					if err != nil {
						log.Println(err.Error())
						return
					}

					setupInterceptorAndClient(token, "178.124.167.214:5300")

					nodeInfoMap = make(map[string]*cardano.SaveStatisticRequest)
					getNodesInfo(true)

					content := container.NewMax()
					title := widget.NewLabel("Component name")

					setTutorial := func(m *MenuField) {
						// Экран каждой ноды
						if fyne.CurrentDevice().IsMobile() {
							child := a.NewWindow(m.Title)
							topWindow = child
							child.SetContent(m.View(topWindow, m.UUID))
							child.Show()
							child.SetOnClosed(func() {
								topWindow = w
							})
							return
						}

						title.SetText(m.Title)

						content.Objects = []fyne.CanvasObject{m.View(w, m.UUID)}
						content.Refresh()
					}

					tutorial := container.NewBorder(
						container.NewVBox(title, widget.NewSeparator()), nil, nil, nil, content,
					)

					var menuNavTree fyne.CanvasObject

					if fyne.CurrentDevice().IsMobile() {
						menuNavTree = makeNav(setTutorial, false)
						w.SetContent(menuNavTree)
					} else {
						menuNavTree = makeNav(setTutorial, true)
						split := container.NewHSplit(menuNavTree, tutorial)

						split.Offset = 0.2
						w.SetContent(split)
					}

					go func() {
						t := time.NewTicker(time.Second * 10)
						for range t.C {
							getNodesInfo(false)
							menuNavTree.Refresh()
						}
					}()

				}),
			),
		),
	)
}

func getNodesInfo(addIndex bool) error {
	resp, err := informClient.GetNodeList()
	if err != nil {
		log.Println(err)
		return err
	}

	cardanoNodes := make(map[string]*cardano.SaveStatisticRequest, 10)
	for _, node := range resp.NodeAuthData {
		switch node.Blockchain {
		case "":
			response, err := cardanoClient.GetStatistic(node.Uuid)
			if err != nil {
				log.Println(err)
				continue
			}

			response.NodeAuthData = node
			cardanoNodes[node.Uuid] = response

			var menuField MenuField
			menuField.Title = node.Ticker
			menuField.UUID = node.Uuid
			menuField.Status = node.Status
			menuField.View = informationScreen

			MenuFields[node.Uuid] = &menuField
			if addIndex {
				MenuIndex[""] = append(MenuIndex[""], node.Uuid)
			}
		}
	}

	nodeInfoMap = cardanoNodes
	return nil
}

var nodeInfoMap map[string]*cardano.SaveStatisticRequest
