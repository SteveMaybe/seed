package seed

func Spinner(spinner int) Seed {
	seed := New()
	seed.tag = "div"
	seed.content = []byte(`<style>.lds-facebook{display:inline-block;position:relative;width:64px;height:64px}.lds-facebook div{display:inline-block;position:absolute;left:6px;width:13px;background:#fff;animation:lds-facebook 1.2s cubic-bezier(0,0.5,0.5,1) infinite}.lds-facebook div:nth-child(1){left:6px;animation-delay:-.24s}.lds-facebook div:nth-child(2){left:26px;animation-delay:-.12s}.lds-facebook div:nth-child(3){left:45px;animation-delay:0}@keyframes lds-facebook{0%{top:6px;height:51px}50%,100%{top:19px;height:26px}}</style>
	<div class="lds-facebook"><div></div><div></div><div></div></div>`)
	return seed
}
