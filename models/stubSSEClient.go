package models

import (
	"fmt"
	"time"
)

type StubSSEClient struct{}

func (c StubSSEClient) Stream(address string, duration time.Duration) (chan Event, error) {
	events := []string{`data: {"instagram_media":{"id":1632042285,"text":"Sabar.. 🦋❤️🥺\nNamaz+Qur'an=Sukoon\nRamzan Mubara k\nFollow 🥺 @ll_ahida_ll\n.\n.\n.\n.\n.\n.\n.\n.\n.\n.\n#ramzan\n#iftari\n#roza\n#follow\n#momin\n#shaikh\n#sanu\n#ahida\n#viral\n#rexplore","link":"","type":"video","location_name":"","likes":4738,"comments":10,"timestamp":1679835328,"post_id":"CqQLjgxuL-g","views":7920,"mtype":3,"tagged_profiles":["ll_ahida_ll","01saniya_sk"],"mentioned_profiles":["ll_ahida_ll"],"thumbnail_url":"https://scontent-iad3-1.cdninstagram.com/v/t51.2885-15/337503390_903679227418190_7673821850889989942_n.jpg?stp=dst-jpg_e15\u0026_nc_ht=scontent-iad3-1.cdninstagram.com\u0026_nc_cat=102\u0026_nc_ohc=6jlyiYRRbX8AX_YM3rQ\u0026edm=AKEQFekBAAAA\u0026ccb=7-5\u0026oh=00_AfD3uMaSu5MKiEPbD6Bj-cE5HSjCpt5fN8Va61yPS6z5UA\u0026oe=6511DB04\u0026_nc_sid=29ddf3","hidden_likes":false,"plays":24688}}`, `data: {"instagram_media":{"id":1697295994,"text":"Does Fort Worth ever cross your mind? ⭐️ In the first pic @tymurraypbr is wearing the belt buckle he won in the bronc riding in 1998 at the @fwssr! Today I got to spend the day with Mimi and his sisters at the @fortworthstockyards! What a fun day exploring and finding out more about the history of Cowtown! From the original brick walkways to the wooden corrals, every inch of the Stockyards tells the true history of Texas’s famous livestock industry. What a special place!  #texas #fortworth","link":"","type":"picture","location_name":"","likes":1153,"comments":21,"timestamp":1695245697,"post_id":"CxbczRSPp20","views":0,"mtype":2,"mentioned_profiles":["tymurraypbr","fwssr","fortworthstockyards"],"thumbnail_url":"https://instagram.fhan3-4.fna.fbcdn.net/v/t51.2885-15/380691182_1040648450423862_6479671342574296590_n.jpg?stp=dst-jpg_e35_p1080x1080\u0026_nc_ht=instagram.fhan3-4.fna.fbcdn.net\u0026_nc_cat=106\u0026_nc_ohc=BR3jctFPgFMAX-spTsm\u0026edm=AKEQFekBAAAA\u0026ccb=7-5\u0026oh=00_AfDaQ6VT1aY_85NQm7hobI3_KjEx9AtDkjac39Pb6OzCOA\u0026oe=65152E5A\u0026_nc_sid=29ddf3","hidden_likes":false,"plays":0}}`, `data: {"pin":{"id":100993011,"title":"DIY Dollar Tree Farmhouse  Fall Garland | Fall Crafts | DIY Fall Décor","description":"Easy DIY dollar store fall garland. Click through for the full fall craft and DIY fall décor tutorial! Dollar tree DIY home décor fall","links":"https://www.mycreativedays.com/diy-dollar-store-fall-garland-craft/","likes":0,"comments":0,"saves":0,"repins":0,"timestamp":1695527768,"post_id":"413557178301350754"}}`, `data: {"pin":{"id":100953162,"title":"","description":" ","links":"https://www.instagram.com/p/BnAJozZjs1J/?igshid=s1ieb0ov3td7","likes":0,"comments":0,"saves":0,"repins":0,"timestamp":1695204160,"post_id":"269019777733782592"}}`, `data: {"instagram_media":{"id":1676771013,"text":"📍Your Dreams","link":"","type":"picture","location_name":"","likes":80,"comments":9,"timestamp":1693703087,"post_id":"Cwtegc2MOiv","views":0,"mtype":2,"thumbnail_url":"https://instagram.fbts3-1.fna.fbcdn.net/v/t51.2885-15/372969355_236042762301145_7200851911473961502_n.jpg?stp=dst-jpg_e35_p1080x1080\u0026_nc_ht=instagram.fbts3-1.fna.fbcdn.net\u0026_nc_cat=101\u0026_nc_ohc=pzFYzLJkYIkAX_AVCNX\u0026edm=AKEQFekBAAAA\u0026ccb=7-5\u0026oh=00_AfCAF9OLOi1tg7UHD4wlQ39V15MHzNAHIsrc-xkG0AeS9A\u0026oe=6515EAC2\u0026_nc_sid=29ddf3","hidden_likes":false,"plays":0}}`, `data: {"instagram_media":{"id":1543498183,"text":"@wanii555  Teri yaadein vo shehar hai 🥀\n\nFollow for more @wanii555\n#writer #poetry #shyari #shayer #wani","link":"","type":"video","location_name":"","likes":1031,"comments":16,"timestamp":1682603355,"post_id":"CrirOEuuqjJ","views":9318,"mtype":3,"tagged_profiles":["azzuuofficial","wanii555"],"mentioned_profiles":["wanii555","wanii555"],"thumbnail_url":"https://instagram.fmel5-1.fna.fbcdn.net/v/t51.2885-15/343391568_242893091574624_1553759238241774016_n.jpg?stp=dst-jpg_e15\u0026_nc_ht=instagram.fmel5-1.fna.fbcdn.net\u0026_nc_cat=100\u0026_nc_ohc=wl3zhM9uix8AX9fnY2j\u0026edm=AKEQFekBAAAA\u0026ccb=7-5\u0026oh=00_AfC2QOopNWwb5ke1oN0fmUvd0DPWNSoxPFGb1PezX0K8qQ\u0026oe=651214C3\u0026_nc_sid=29ddf3","hidden_likes":false,"plays":25045}}`, `data: {"pin":{"id":100868824,"title":"","description":" ","links":"https://www.savoringthegood.com/masterbuilt-smoker-recipes/","likes":0,"comments":0,"saves":0,"repins":0,"timestamp":1694530994,"post_id":"272327108711925220"}}`, `data: {"tiktok_video":{"id":277682097,"name":"faz o vuck vuck.. ","thumbnail_url":"https://p16-sign-va.tiktokcdn.com/obj/tos-maliva-p-0068/osyVkuIZ5AiEGAfIKg4IRpmBChCzodNUu83zhm?x-expires=1695729600\u0026x-signature=%2FjiFToGv%2Bu8IBppJ1dbEEY8l5FU%3D","link":"https://www.tiktok.com/@kaisaborges/video/7280313537287687430","comments":2,"likes":36,"timestamp":1695080091,"post_id":"7280313537287687430","plays":278,"shares":0}}`, `data: {"instagram_media":{"id":1699696135,"text":"18 months of grey hair!\n\nHaven't looked back 👵�\\n\n#silverhair #silversisters #greyhair","link":"","type":"picture","location_name":"","likes":136,"comments":47,"timestamp":1695108328,"post_id":"CxXWyrYMDv5","views":0,"mtype":2,"thumbnail_url":"https://instagram.fsgn5-6.fna.fbcdn.net/v/t51.2885-15/380279725_165258003288080_2907648033519837892_n.heic?stp=dst-jpg_e35_s1080x1080\u0026_nc_ht=instagram.fsgn5-6.fna.fbcdn.net\u0026_nc_cat=108\u0026_nc_ohc=rtiohIQQp68AX_jzjxt\u0026edm=AKEQFekBAAAA\u0026ccb=7-5\u0026oh=00_AfCLoRireW4mEN06DVnE3wl_v17kWeGz0J60eXoE3I_vgg\u0026oe=65152E79\u0026_nc_sid=29ddf3","hidden_likes":false,"plays":0}}`}
	eventChan := make(chan Event)

	go func() {
		continueLoop := true
		loopCounter := 0
		time.AfterFunc(duration, func() { //close channel after duration
			close(eventChan)
			continueLoop = false
		})
		for now := range time.Tick(250 * time.Millisecond) { //send 4 events/second
			loopCounter++
			fmt.Println(now)
			if !continueLoop { //Security before sending in a closed channel
				return
			}
			event, err := NewEvent(events[loopCounter%len(events)])
			if err != nil {
				fmt.Println(err)
				continue
			}
			eventChan <- event
		}
	}()

	return eventChan, nil

}
