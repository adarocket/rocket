package main

import (
	"fmt"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/adarocket/proto"
)

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

func informationScreen(w fyne.Window, uuid string) fyne.CanvasObject {
	intro := widget.NewLabel("No information about this node")
	intro.Wrapping = fyne.TextWrapWord

	resp := nodeInfoMap[uuid]

	var items []fyne.CanvasObject
	item := new(widget.Form)

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

					setupInterceptorAndClient(token)

					nodeInfoMap = make(map[string]*proto.SaveStatisticRequest)
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
		return nil
	}

	for _, node := range resp.NodeAuthData {
		var menuField MenuField
		menuField.Title = node.Ticker
		menuField.UUID = node.Uuid
		menuField.Status = node.Status
		menuField.View = informationScreen

		MenuFields[node.Uuid] = &menuField
		if addIndex {
			MenuIndex[""] = append(MenuIndex[""], node.Uuid)
		}

		resp, err := informClient.GetStatistic(node.Uuid)
		if err != nil {
			log.Println(err)
			fyne.LogError("Error!!!", err)
			continue
		}
		nodeInfoMap[node.Uuid] = resp
	}

	return nil
}

var nodeInfoMap map[string]*proto.SaveStatisticRequest

// func informationScreen(w fyne.Window, uuid string) fyne.CanvasObject {
// 	// fyne.LogError("informationScreen uuid"+uuid, errors.New("informationScreen uuid"+uuid))

// 	intro := widget.NewLabel("No information about this node")
// 	intro.Wrapping = fyne.TextWrapWord

// 	resp := nodeInfoMap[uuid]
// 	// resp, err := informClient.GetStatistic(uuid)
// 	// if err != nil {
// 	// 	log.Println(err)
// 	// 	fyne.LogError("Error!!!", err)
// 	// 	return container.NewVBox(intro)
// 	// }

// 	var items []*widget.AccordionItem
// 	if resp.Statistic.NodeBasicData != nil {
// 		item := new(widget.AccordionItem)
// 		item = widget.NewAccordionItem("Node Basic Data",
// 			widget.NewForm(
// 				widget.NewFormItem("Ticker", widget.NewLabel(resp.Statistic.NodeBasicData.Ticker)),
// 				widget.NewFormItem("Type", widget.NewLabel(resp.Statistic.NodeBasicData.Type)),
// 				widget.NewFormItem("Location", widget.NewLabel(resp.Statistic.NodeBasicData.Location)),
// 				widget.NewFormItem("Node version", widget.NewLabel(resp.Statistic.NodeBasicData.NodeVersion)),
// 			),
// 		)
// 		items = append(items, item)
// 	}

// 	if resp.Statistic.ServerBasicData != nil {
// 		item := new(widget.AccordionItem)
// 		item = widget.NewAccordionItem("Server Basic Data",
// 			widget.NewForm(
// 				widget.NewFormItem("IPv4", widget.NewLabel(resp.Statistic.ServerBasicData.Ipv4)),
// 				widget.NewFormItem("IPv6", widget.NewLabel(resp.Statistic.ServerBasicData.Ipv6)),
// 				widget.NewFormItem("Linux name", widget.NewLabel(resp.Statistic.ServerBasicData.LinuxName)),
// 				widget.NewFormItem("Linux version", widget.NewLabel(resp.Statistic.ServerBasicData.LinuxVersion)),
// 			),
// 		)
// 		items = append(items, item)
// 	}

// 	if resp.Statistic.Online != nil {
// 		item := new(widget.AccordionItem)
// 		item = widget.NewAccordionItem("Online",
// 			widget.NewForm(
// 				widget.NewFormItem("Since start", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.Online.SinceStart))),
// 				widget.NewFormItem("Pings", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.Online.Pings))),
// 				widget.NewFormItem("Node active", widget.NewLabel(fmt.Sprintf("%t", resp.Statistic.Online.NodeActive))),
// 				widget.NewFormItem("Node active pings", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.Online.NodeActivePings))),
// 			),
// 		)
// 		items = append(items, item)
// 	}

// 	if resp.Statistic.MemoryState != nil {
// 		item := new(widget.AccordionItem)
// 		item = widget.NewAccordionItem("Memory",
// 			widget.NewForm(
// 				widget.NewFormItem("Total", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.MemoryState.Total))),
// 				widget.NewFormItem("Used", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.MemoryState.Used))),
// 				widget.NewFormItem("Buffers", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.MemoryState.Buffers))),
// 				widget.NewFormItem("Cached", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.MemoryState.Cached))),
// 				widget.NewFormItem("Free", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.MemoryState.Free))),
// 				widget.NewFormItem("Available", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.MemoryState.Available))),
// 				widget.NewFormItem("Active", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.MemoryState.Active))),
// 				widget.NewFormItem("Inactive", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.MemoryState.Inactive))),
// 				widget.NewFormItem("Swap Total", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.MemoryState.SwapTotal))),
// 				widget.NewFormItem("Swap Used", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.MemoryState.SwapUsed))),
// 				widget.NewFormItem("Swap Cached", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.MemoryState.SwapCached))),
// 				widget.NewFormItem("Swap Free", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.MemoryState.SwapFree))),
// 				widget.NewFormItem("Mem Available Enabled", widget.NewLabel(fmt.Sprintf("%t", resp.Statistic.MemoryState.MemAvailableEnabled))),
// 			),
// 		)
// 		items = append(items, item)
// 	}

// 	if resp.Statistic.CpuState != nil {
// 		item := new(widget.AccordionItem)
// 		item = widget.NewAccordionItem("CPU state",
// 			widget.NewForm(
// 				widget.NewFormItem("CPU Qty", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.CpuState.CpuQty))),
// 				widget.NewFormItem("Average workload", widget.NewLabel(fmt.Sprintf("%f", resp.Statistic.CpuState.AverageWorkload))),
// 			),
// 		)
// 		items = append(items, item)
// 	}

// 	if resp.Statistic.Epoch != nil {
// 		item := new(widget.AccordionItem)
// 		item = widget.NewAccordionItem("Epoch",
// 			widget.NewForm(
// 				widget.NewFormItem("Epoch number", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.Epoch.EpochNumber))),
// 			),
// 		)
// 		items = append(items, item)
// 	}

// 	if resp.Statistic.NodeState != nil {
// 		item := new(widget.AccordionItem)
// 		item = widget.NewAccordionItem("Node State",
// 			widget.NewForm(
// 				widget.NewFormItem("Tip diff", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.NodeState.TipDiff))),
// 				widget.NewFormItem("Density", widget.NewLabel(fmt.Sprintf("%f", resp.Statistic.NodeState.Density)))),
// 		)
// 		items = append(items, item)
// 	}

// 	if resp.Statistic.NodePerformance != nil {
// 		item := new(widget.AccordionItem)
// 		item = widget.NewAccordionItem("Node Performance",
// 			widget.NewForm(
// 				widget.NewFormItem("Processed Tx", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.NodePerformance.ProcessedTx))),
// 				widget.NewFormItem("Peers In", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.NodePerformance.PeersIn))),
// 				widget.NewFormItem("Peers Out", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.NodePerformance.PeersOut)))),
// 		)
// 		items = append(items, item)
// 	}

// 	if resp.Statistic.KesData != nil {
// 		item := new(widget.AccordionItem)
// 		item = widget.NewAccordionItem("KES Data",
// 			widget.NewForm(
// 				widget.NewFormItem("KES current", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.KesData.KesCurrent))),
// 				widget.NewFormItem("KES remaining", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.KesData.KesRemaining))),
// 				widget.NewFormItem("KES exp date", widget.NewLabel(resp.Statistic.KesData.KesExpDate))),
// 		)
// 		items = append(items, item)
// 	}

// 	if resp.Statistic.Blocks != nil {
// 		item := new(widget.AccordionItem)
// 		item = widget.NewAccordionItem("Blocks",
// 			widget.NewForm(
// 				widget.NewFormItem("Block leader", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.Blocks.BlockLeader))),
// 				widget.NewFormItem("Block adopted", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.Blocks.BlockAdopted))),
// 				widget.NewFormItem("Block invalid", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.Blocks.BlockInvalid)))),
// 		)
// 		items = append(items, item)
// 	}

// 	if resp.Statistic.Updates != nil {
// 		item := new(widget.AccordionItem)
// 		item = widget.NewAccordionItem("Updates",
// 			widget.NewForm(
// 				widget.NewFormItem("Informer actual", widget.NewLabel(resp.Statistic.Updates.InformerActual)),
// 				widget.NewFormItem("Informer available", widget.NewLabel(resp.Statistic.Updates.InformerAvailable)),
// 				widget.NewFormItem("Updater actual", widget.NewLabel(resp.Statistic.Updates.UpdaterActual)),
// 				widget.NewFormItem("Updater available", widget.NewLabel(resp.Statistic.Updates.UpdaterAvailable)),
// 				widget.NewFormItem("Packages available", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.Updates.PackagesAvailable))),
// 			),
// 		)
// 		items = append(items, item)
// 	}

// 	if resp.Statistic.Security != nil {
// 		item := new(widget.AccordionItem)
// 		item = widget.NewAccordionItem("Security",
// 			widget.NewForm(
// 				widget.NewFormItem("SSH Attack Attempts", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.Security.SshAttackAttempts))),
// 				widget.NewFormItem("Security Packages Available", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.Security.SecurityPackagesAvailable))),
// 			),
// 		)
// 		items = append(items, item)
// 	}

// 	if resp.Statistic.StakeInfo != nil {
// 		item := new(widget.AccordionItem)
// 		item = widget.NewAccordionItem("StakeInfo",
// 			widget.NewForm(
// 				widget.NewFormItem("Live stake", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.StakeInfo.LiveStake))),
// 				widget.NewFormItem("Active stake", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.StakeInfo.ActiveStake))),
// 				widget.NewFormItem("Pledge", widget.NewLabel(fmt.Sprintf("%d", resp.Statistic.StakeInfo.Pledge))),
// 			),
// 		)
// 		items = append(items, item)
// 	}

// 	if len(items) > 0 {
// 		return container.NewVScroll(container.NewVBox(widget.NewAccordion(items...)))
// 	}

// 	// fyne.LogError("Items les then 0", errors.New("Items les then 0"))
// 	return container.NewVBox(intro)
// }