package structs
type LoLChampionKled struct{
	Type string`json:"type"`
	Format string`json:"format"`
	Version string`json:"version"`
	Data struct{
		Kled struct{
			Id string`json:"id"`
			Key string`json:"key"`
			Name string`json:"name"`
			Title string`json:"title"`
			Image struct{
				Full string`json:"full"`
				Sprite string`json:"sprite"`
				Group string`json:"group"`
				X int`json:"x"`
				Y int`json:"y"`
				W int`json:"w"`
				H int`json:"h"`
			}`json:"image"`
			Skins struct{
				Skins0 struct{
					Id string`json:"id"`
					Num int`json:"num"`
					Name string`json:"name"`
					Chromas bool`json:"chromas"`
				}`json:"skins0"`
				Skins1 struct{
					Id string`json:"id"`
					Num int`json:"num"`
					Name string`json:"name"`
					Chromas bool`json:"chromas"`
				}`json:"skins1"`
				Skins2 struct{
					Id string`json:"id"`
					Num int`json:"num"`
					Name string`json:"name"`
					Chromas bool`json:"chromas"`
				}`json:"skins2"`
				Skins3 struct{
					Id string`json:"id"`
					Num int`json:"num"`
					Name string`json:"name"`
					Chromas bool`json:"chromas"`
				}`json:"skins3"`
				Skins4 struct{
					Id string`json:"id"`
					Num int`json:"num"`
					Name string`json:"name"`
					Chromas bool`json:"chromas"`
				}`json:"skins4"`
			}`json:"skins"`
			Lore string`json:"lore"`
			Blurb string`json:"blurb"`
			Allytips struct{
				Allytips0 string`json:"allytips0"`
				Allytips1 string`json:"allytips1"`
				Allytips2 string`json:"allytips2"`
			}`json:"allytips"`
			Enemytips struct{
				Enemytips0 string`json:"enemytips0"`
				Enemytips1 string`json:"enemytips1"`
				Enemytips2 string`json:"enemytips2"`
			}`json:"enemytips"`
			Tags struct{
				Tags0 string`json:"tags0"`
				Tags1 string`json:"tags1"`
			}`json:"tags"`
			Partype string`json:"partype"`
			Info struct{
				Attack int`json:"attack"`
				Defense int`json:"defense"`
				Magic int`json:"magic"`
				Difficulty int`json:"difficulty"`
			}`json:"info"`
			Stats struct{
				Hp int`json:"hp"`
				Hpperlevel int`json:"hpperlevel"`
				Mp int`json:"mp"`
				Mpperlevel int`json:"mpperlevel"`
				Movespeed int`json:"movespeed"`
				Armor int`json:"armor"`
				Armorperlevel float32`json:"armorperlevel"`
				Spellblock int`json:"spellblock"`
				Spellblockperlevel float32`json:"spellblockperlevel"`
				Attackrange int`json:"attackrange"`
				Hpregen int`json:"hpregen"`
				Hpregenperlevel float32`json:"hpregenperlevel"`
				Mpregen int`json:"mpregen"`
				Mpregenperlevel int`json:"mpregenperlevel"`
				Crit int`json:"crit"`
				Critperlevel int`json:"critperlevel"`
				Attackdamage int`json:"attackdamage"`
				Attackdamageperlevel float32`json:"attackdamageperlevel"`
				Attackspeedperlevel float32`json:"attackspeedperlevel"`
				Attackspeed float32`json:"attackspeed"`
			}`json:"stats"`
			Spells struct{
				Spells0 struct{
					Id string`json:"id"`
					Name string`json:"name"`
					Description string`json:"description"`
					Tooltip string`json:"tooltip"`
					Leveltip struct{
						Label struct{
							Label0 string`json:"label0"`
							Label1 string`json:"label1"`
							Label2 string`json:"label2"`
							Label3 string`json:"label3"`
						}`json:"label"`
						Effect struct{
							Effect0 string`json:"effect0"`
							Effect1 string`json:"effect1"`
							Effect2 string`json:"effect2"`
							Effect3 string`json:"effect3"`
						}`json:"effect"`
					}`json:"leveltip"`
					Maxrank int`json:"maxrank"`
					Cooldown struct{
						Cooldown0 int`json:"cooldown0"`
						Cooldown1 int`json:"cooldown1"`
						Cooldown2 int`json:"cooldown2"`
						Cooldown3 int`json:"cooldown3"`
						Cooldown4 int`json:"cooldown4"`
					}`json:"cooldown"`
					CooldownBurn string`json:"cooldownBurn"`
					Cost struct{
						Cost0 int`json:"cost0"`
						Cost1 int`json:"cost1"`
						Cost2 int`json:"cost2"`
						Cost3 int`json:"cost3"`
						Cost4 int`json:"cost4"`
					}`json:"cost"`
					CostBurn string`json:"costBurn"`
					Datavalues struct{
					}`json:"datavalues"`
					Effect struct{
						Effect0 nil`json:"effect0"`
						Effect1 struct{
							Effect10 int`json:"effect10"`
							Effect11 int`json:"effect11"`
							Effect12 int`json:"effect12"`
							Effect13 int`json:"effect13"`
							Effect14 int`json:"effect14"`
						}`json:"effect1"`
						Effect2 struct{
							Effect20 float32`json:"effect20"`
							Effect21 float32`json:"effect21"`
							Effect22 float32`json:"effect22"`
							Effect23 float32`json:"effect23"`
							Effect24 float32`json:"effect24"`
						}`json:"effect2"`
						Effect3 struct{
							Effect30 float32`json:"effect30"`
							Effect31 float32`json:"effect31"`
							Effect32 float32`json:"effect32"`
							Effect33 float32`json:"effect33"`
							Effect34 float32`json:"effect34"`
						}`json:"effect3"`
						Effect4 struct{
							Effect40 int`json:"effect40"`
							Effect41 int`json:"effect41"`
							Effect42 int`json:"effect42"`
							Effect43 int`json:"effect43"`
							Effect44 int`json:"effect44"`
						}`json:"effect4"`
						Effect5 struct{
							Effect50 int`json:"effect50"`
							Effect51 int`json:"effect51"`
							Effect52 int`json:"effect52"`
							Effect53 int`json:"effect53"`
							Effect54 int`json:"effect54"`
						}`json:"effect5"`
						Effect6 struct{
							Effect60 int`json:"effect60"`
							Effect61 float32`json:"effect61"`
							Effect62 int`json:"effect62"`
							Effect63 float32`json:"effect63"`
							Effect64 int`json:"effect64"`
						}`json:"effect6"`
						Effect7 struct{
							Effect70 float32`json:"effect70"`
							Effect71 float32`json:"effect71"`
							Effect72 float32`json:"effect72"`
							Effect73 float32`json:"effect73"`
							Effect74 float32`json:"effect74"`
						}`json:"effect7"`
						Effect8 struct{
							Effect80 int`json:"effect80"`
							Effect81 int`json:"effect81"`
							Effect82 int`json:"effect82"`
							Effect83 int`json:"effect83"`
							Effect84 int`json:"effect84"`
						}`json:"effect8"`
						Effect9 struct{
							Effect90 int`json:"effect90"`
							Effect91 int`json:"effect91"`
							Effect92 int`json:"effect92"`
							Effect93 int`json:"effect93"`
							Effect94 int`json:"effect94"`
						}`json:"effect9"`
						Effect10 struct{
							Effect100 int`json:"effect100"`
							Effect101 int`json:"effect101"`
							Effect102 int`json:"effect102"`
							Effect103 int`json:"effect103"`
							Effect104 int`json:"effect104"`
						}`json:"effect10"`
					}`json:"effect"`
					EffectBurn struct{
						EffectBurn0 nil`json:"effectBurn0"`
						EffectBurn1 string`json:"effectBurn1"`
						EffectBurn2 string`json:"effectBurn2"`
						EffectBurn3 string`json:"effectBurn3"`
						EffectBurn4 string`json:"effectBurn4"`
						EffectBurn5 string`json:"effectBurn5"`
						EffectBurn6 string`json:"effectBurn6"`
						EffectBurn7 string`json:"effectBurn7"`
						EffectBurn8 string`json:"effectBurn8"`
						EffectBurn9 string`json:"effectBurn9"`
						EffectBurn10 string`json:"effectBurn10"`
					}`json:"effectBurn"`
					Vars struct{
					}`json:"vars"`
					CostType string`json:"costType"`
					Maxammo string`json:"maxammo"`
					Range struct{
						Range0 int`json:"range0"`
						Range1 int`json:"range1"`
						Range2 int`json:"range2"`
						Range3 int`json:"range3"`
						Range4 int`json:"range4"`
					}`json:"range"`
					RangeBurn string`json:"rangeBurn"`
					Image struct{
						Full string`json:"full"`
						Sprite string`json:"sprite"`
						Group string`json:"group"`
						X int`json:"x"`
						Y int`json:"y"`
						W int`json:"w"`
						H int`json:"h"`
					}`json:"image"`
					Resource string`json:"resource"`
				}`json:"spells0"`
				Spells1 struct{
					Id string`json:"id"`
					Name string`json:"name"`
					Description string`json:"description"`
					Tooltip string`json:"tooltip"`
					Leveltip struct{
						Label struct{
							Label0 string`json:"label0"`
							Label1 string`json:"label1"`
							Label2 string`json:"label2"`
						}`json:"label"`
						Effect struct{
							Effect0 string`json:"effect0"`
							Effect1 string`json:"effect1"`
							Effect2 string`json:"effect2"`
						}`json:"effect"`
					}`json:"leveltip"`
					Maxrank int`json:"maxrank"`
					Cooldown struct{
						Cooldown0 int`json:"cooldown0"`
						Cooldown1 float32`json:"cooldown1"`
						Cooldown2 int`json:"cooldown2"`
						Cooldown3 float32`json:"cooldown3"`
						Cooldown4 int`json:"cooldown4"`
					}`json:"cooldown"`
					CooldownBurn string`json:"cooldownBurn"`
					Cost struct{
						Cost0 int`json:"cost0"`
						Cost1 int`json:"cost1"`
						Cost2 int`json:"cost2"`
						Cost3 int`json:"cost3"`
						Cost4 int`json:"cost4"`
					}`json:"cost"`
					CostBurn string`json:"costBurn"`
					Datavalues struct{
					}`json:"datavalues"`
					Effect struct{
						Effect0 nil`json:"effect0"`
						Effect1 struct{
							Effect10 float32`json:"effect10"`
							Effect11 int`json:"effect11"`
							Effect12 float32`json:"effect12"`
							Effect13 int`json:"effect13"`
							Effect14 float32`json:"effect14"`
						}`json:"effect1"`
						Effect2 struct{
							Effect20 int`json:"effect20"`
							Effect21 int`json:"effect21"`
							Effect22 int`json:"effect22"`
							Effect23 int`json:"effect23"`
							Effect24 int`json:"effect24"`
						}`json:"effect2"`
						Effect3 struct{
							Effect30 int`json:"effect30"`
							Effect31 int`json:"effect31"`
							Effect32 int`json:"effect32"`
							Effect33 int`json:"effect33"`
							Effect34 int`json:"effect34"`
						}`json:"effect3"`
						Effect4 struct{
							Effect40 int`json:"effect40"`
							Effect41 int`json:"effect41"`
							Effect42 int`json:"effect42"`
							Effect43 int`json:"effect43"`
							Effect44 int`json:"effect44"`
						}`json:"effect4"`
						Effect5 struct{
							Effect50 int`json:"effect50"`
							Effect51 int`json:"effect51"`
							Effect52 int`json:"effect52"`
							Effect53 int`json:"effect53"`
							Effect54 int`json:"effect54"`
						}`json:"effect5"`
						Effect6 struct{
							Effect60 int`json:"effect60"`
							Effect61 int`json:"effect61"`
							Effect62 int`json:"effect62"`
							Effect63 int`json:"effect63"`
							Effect64 int`json:"effect64"`
						}`json:"effect6"`
						Effect7 struct{
							Effect70 int`json:"effect70"`
							Effect71 int`json:"effect71"`
							Effect72 int`json:"effect72"`
							Effect73 int`json:"effect73"`
							Effect74 int`json:"effect74"`
						}`json:"effect7"`
						Effect8 struct{
							Effect80 int`json:"effect80"`
							Effect81 int`json:"effect81"`
							Effect82 int`json:"effect82"`
							Effect83 int`json:"effect83"`
							Effect84 int`json:"effect84"`
						}`json:"effect8"`
						Effect9 struct{
							Effect90 int`json:"effect90"`
							Effect91 int`json:"effect91"`
							Effect92 int`json:"effect92"`
							Effect93 int`json:"effect93"`
							Effect94 int`json:"effect94"`
						}`json:"effect9"`
						Effect10 struct{
							Effect100 int`json:"effect100"`
							Effect101 int`json:"effect101"`
							Effect102 int`json:"effect102"`
							Effect103 int`json:"effect103"`
							Effect104 int`json:"effect104"`
						}`json:"effect10"`
					}`json:"effect"`
					EffectBurn struct{
						EffectBurn0 nil`json:"effectBurn0"`
						EffectBurn1 string`json:"effectBurn1"`
						EffectBurn2 string`json:"effectBurn2"`
						EffectBurn3 string`json:"effectBurn3"`
						EffectBurn4 string`json:"effectBurn4"`
						EffectBurn5 string`json:"effectBurn5"`
						EffectBurn6 string`json:"effectBurn6"`
						EffectBurn7 string`json:"effectBurn7"`
						EffectBurn8 string`json:"effectBurn8"`
						EffectBurn9 string`json:"effectBurn9"`
						EffectBurn10 string`json:"effectBurn10"`
					}`json:"effectBurn"`
					Vars struct{
					}`json:"vars"`
					CostType string`json:"costType"`
					Maxammo string`json:"maxammo"`
					Range struct{
						Range0 int`json:"range0"`
						Range1 int`json:"range1"`
						Range2 int`json:"range2"`
						Range3 int`json:"range3"`
						Range4 int`json:"range4"`
					}`json:"range"`
					RangeBurn string`json:"rangeBurn"`
					Image struct{
						Full string`json:"full"`
						Sprite string`json:"sprite"`
						Group string`json:"group"`
						X int`json:"x"`
						Y int`json:"y"`
						W int`json:"w"`
						H int`json:"h"`
					}`json:"image"`
					Resource string`json:"resource"`
				}`json:"spells1"`
				Spells2 struct{
					Id string`json:"id"`
					Name string`json:"name"`
					Description string`json:"description"`
					Tooltip string`json:"tooltip"`
					Leveltip struct{
						Label struct{
							Label0 string`json:"label0"`
							Label1 string`json:"label1"`
						}`json:"label"`
						Effect struct{
							Effect0 string`json:"effect0"`
							Effect1 string`json:"effect1"`
						}`json:"effect"`
					}`json:"leveltip"`
					Maxrank int`json:"maxrank"`
					Cooldown struct{
						Cooldown0 int`json:"cooldown0"`
						Cooldown1 int`json:"cooldown1"`
						Cooldown2 int`json:"cooldown2"`
						Cooldown3 int`json:"cooldown3"`
						Cooldown4 int`json:"cooldown4"`
					}`json:"cooldown"`
					CooldownBurn string`json:"cooldownBurn"`
					Cost struct{
						Cost0 int`json:"cost0"`
						Cost1 int`json:"cost1"`
						Cost2 int`json:"cost2"`
						Cost3 int`json:"cost3"`
						Cost4 int`json:"cost4"`
					}`json:"cost"`
					CostBurn string`json:"costBurn"`
					Datavalues struct{
					}`json:"datavalues"`
					Effect struct{
						Effect0 nil`json:"effect0"`
						Effect1 struct{
							Effect10 int`json:"effect10"`
							Effect11 int`json:"effect11"`
							Effect12 int`json:"effect12"`
							Effect13 int`json:"effect13"`
							Effect14 int`json:"effect14"`
						}`json:"effect1"`
						Effect2 struct{
							Effect20 int`json:"effect20"`
							Effect21 int`json:"effect21"`
							Effect22 int`json:"effect22"`
							Effect23 int`json:"effect23"`
							Effect24 int`json:"effect24"`
						}`json:"effect2"`
						Effect3 struct{
							Effect30 int`json:"effect30"`
							Effect31 int`json:"effect31"`
							Effect32 int`json:"effect32"`
							Effect33 int`json:"effect33"`
							Effect34 int`json:"effect34"`
						}`json:"effect3"`
						Effect4 struct{
							Effect40 int`json:"effect40"`
							Effect41 int`json:"effect41"`
							Effect42 int`json:"effect42"`
							Effect43 int`json:"effect43"`
							Effect44 int`json:"effect44"`
						}`json:"effect4"`
						Effect5 struct{
							Effect50 int`json:"effect50"`
							Effect51 int`json:"effect51"`
							Effect52 int`json:"effect52"`
							Effect53 int`json:"effect53"`
							Effect54 int`json:"effect54"`
						}`json:"effect5"`
						Effect6 struct{
							Effect60 int`json:"effect60"`
							Effect61 int`json:"effect61"`
							Effect62 int`json:"effect62"`
							Effect63 int`json:"effect63"`
							Effect64 int`json:"effect64"`
						}`json:"effect6"`
						Effect7 struct{
							Effect70 int`json:"effect70"`
							Effect71 int`json:"effect71"`
							Effect72 int`json:"effect72"`
							Effect73 int`json:"effect73"`
							Effect74 int`json:"effect74"`
						}`json:"effect7"`
						Effect8 struct{
							Effect80 int`json:"effect80"`
							Effect81 int`json:"effect81"`
							Effect82 int`json:"effect82"`
							Effect83 int`json:"effect83"`
							Effect84 int`json:"effect84"`
						}`json:"effect8"`
						Effect9 struct{
							Effect90 int`json:"effect90"`
							Effect91 int`json:"effect91"`
							Effect92 int`json:"effect92"`
							Effect93 int`json:"effect93"`
							Effect94 int`json:"effect94"`
						}`json:"effect9"`
						Effect10 struct{
							Effect100 int`json:"effect100"`
							Effect101 int`json:"effect101"`
							Effect102 int`json:"effect102"`
							Effect103 int`json:"effect103"`
							Effect104 int`json:"effect104"`
						}`json:"effect10"`
					}`json:"effect"`
					EffectBurn struct{
						EffectBurn0 nil`json:"effectBurn0"`
						EffectBurn1 string`json:"effectBurn1"`
						EffectBurn2 string`json:"effectBurn2"`
						EffectBurn3 string`json:"effectBurn3"`
						EffectBurn4 string`json:"effectBurn4"`
						EffectBurn5 string`json:"effectBurn5"`
						EffectBurn6 string`json:"effectBurn6"`
						EffectBurn7 string`json:"effectBurn7"`
						EffectBurn8 string`json:"effectBurn8"`
						EffectBurn9 string`json:"effectBurn9"`
						EffectBurn10 string`json:"effectBurn10"`
					}`json:"effectBurn"`
					Vars struct{
					}`json:"vars"`
					CostType string`json:"costType"`
					Maxammo string`json:"maxammo"`
					Range struct{
						Range0 int`json:"range0"`
						Range1 int`json:"range1"`
						Range2 int`json:"range2"`
						Range3 int`json:"range3"`
						Range4 int`json:"range4"`
					}`json:"range"`
					RangeBurn string`json:"rangeBurn"`
					Image struct{
						Full string`json:"full"`
						Sprite string`json:"sprite"`
						Group string`json:"group"`
						X int`json:"x"`
						Y int`json:"y"`
						W int`json:"w"`
						H int`json:"h"`
					}`json:"image"`
					Resource string`json:"resource"`
				}`json:"spells2"`
				Spells3 struct{
					Id string`json:"id"`
					Name string`json:"name"`
					Description string`json:"description"`
					Tooltip string`json:"tooltip"`
					Leveltip struct{
						Label struct{
							Label0 string`json:"label0"`
							Label1 string`json:"label1"`
							Label2 string`json:"label2"`
							Label3 string`json:"label3"`
						}`json:"label"`
						Effect struct{
							Effect0 string`json:"effect0"`
							Effect1 string`json:"effect1"`
							Effect2 string`json:"effect2"`
							Effect3 string`json:"effect3"`
						}`json:"effect"`
					}`json:"leveltip"`
					Maxrank int`json:"maxrank"`
					Cooldown struct{
						Cooldown0 int`json:"cooldown0"`
						Cooldown1 int`json:"cooldown1"`
						Cooldown2 int`json:"cooldown2"`
					}`json:"cooldown"`
					CooldownBurn string`json:"cooldownBurn"`
					Cost struct{
						Cost0 int`json:"cost0"`
						Cost1 int`json:"cost1"`
						Cost2 int`json:"cost2"`
					}`json:"cost"`
					CostBurn string`json:"costBurn"`
					Datavalues struct{
					}`json:"datavalues"`
					Effect struct{
						Effect0 nil`json:"effect0"`
						Effect1 struct{
							Effect10 int`json:"effect10"`
							Effect11 int`json:"effect11"`
							Effect12 int`json:"effect12"`
						}`json:"effect1"`
						Effect2 struct{
							Effect20 int`json:"effect20"`
							Effect21 int`json:"effect21"`
							Effect22 int`json:"effect22"`
						}`json:"effect2"`
						Effect3 struct{
							Effect30 int`json:"effect30"`
							Effect31 int`json:"effect31"`
							Effect32 int`json:"effect32"`
						}`json:"effect3"`
						Effect4 struct{
							Effect40 int`json:"effect40"`
							Effect41 int`json:"effect41"`
							Effect42 int`json:"effect42"`
						}`json:"effect4"`
						Effect5 struct{
							Effect50 int`json:"effect50"`
							Effect51 int`json:"effect51"`
							Effect52 int`json:"effect52"`
						}`json:"effect5"`
						Effect6 struct{
							Effect60 int`json:"effect60"`
							Effect61 int`json:"effect61"`
							Effect62 int`json:"effect62"`
						}`json:"effect6"`
						Effect7 struct{
							Effect70 int`json:"effect70"`
							Effect71 int`json:"effect71"`
							Effect72 int`json:"effect72"`
						}`json:"effect7"`
						Effect8 struct{
							Effect80 int`json:"effect80"`
							Effect81 int`json:"effect81"`
							Effect82 int`json:"effect82"`
						}`json:"effect8"`
						Effect9 struct{
							Effect90 int`json:"effect90"`
							Effect91 int`json:"effect91"`
							Effect92 int`json:"effect92"`
						}`json:"effect9"`
						Effect10 struct{
							Effect100 int`json:"effect100"`
							Effect101 int`json:"effect101"`
							Effect102 int`json:"effect102"`
						}`json:"effect10"`
					}`json:"effect"`
					EffectBurn struct{
						EffectBurn0 nil`json:"effectBurn0"`
						EffectBurn1 string`json:"effectBurn1"`
						EffectBurn2 string`json:"effectBurn2"`
						EffectBurn3 string`json:"effectBurn3"`
						EffectBurn4 string`json:"effectBurn4"`
						EffectBurn5 string`json:"effectBurn5"`
						EffectBurn6 string`json:"effectBurn6"`
						EffectBurn7 string`json:"effectBurn7"`
						EffectBurn8 string`json:"effectBurn8"`
						EffectBurn9 string`json:"effectBurn9"`
						EffectBurn10 string`json:"effectBurn10"`
					}`json:"effectBurn"`
					Vars struct{
					}`json:"vars"`
					CostType string`json:"costType"`
					Maxammo string`json:"maxammo"`
					Range struct{
						Range0 int`json:"range0"`
						Range1 int`json:"range1"`
						Range2 int`json:"range2"`
					}`json:"range"`
					RangeBurn string`json:"rangeBurn"`
					Image struct{
						Full string`json:"full"`
						Sprite string`json:"sprite"`
						Group string`json:"group"`
						X int`json:"x"`
						Y int`json:"y"`
						W int`json:"w"`
						H int`json:"h"`
					}`json:"image"`
					Resource string`json:"resource"`
				}`json:"spells3"`
			}`json:"spells"`
			Passive struct{
				Name string`json:"name"`
				Description string`json:"description"`
				Image struct{
					Full string`json:"full"`
					Sprite string`json:"sprite"`
					Group string`json:"group"`
					X int`json:"x"`
					Y int`json:"y"`
					W int`json:"w"`
					H int`json:"h"`
				}`json:"image"`
			}`json:"passive"`
			Recommended struct{
			}`json:"recommended"`
		}`json:"Kled"`
	}`json:"data"`
}
