package dscan

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/ZmeisIncorporated/pochven-map/pkg/types"

)

func TestBasicDscan(t *testing.T) {
	rawDscan := `
47513	Rakapas - Copa Cabana	'Draccous' Fortizar	5,3 AU
37604	Feedot	Apostle	-
35835	Rakapas - Reproc Ice - Composite React	Athanor	5,4 AU
35835	Rakapas - Reproc Ore+Moon - Hybrid+Bio	Athanor	5,3 AU
35826	Rakapas - CSAA	Azbel	5,3 AU
1529	Rakapas V - Home Guard Assembly Plant	Caldari Administrative Station	5,3 AU
1530	Rakapas II - State Protectorate Logistic Support	Caldari Research Station	2,5 AU
35833	Rakapas - Fox Hole	Fortizar	2,5 AU
37606	Needs Fit	Lif	-
37606	LIF	Lif	-
44996	reddog	Marshal	-
35825	Rakapas - Copying-Research	Raitaru	5,3 AU
35825	Rakapas - Manufacturing	Raitaru	5,3 AU
35825	Rakapas - Invention-Manufacturing	Raitaru	5,3 AU
22428	救世啊	Redeemer	-
16	Iwisoda	Stargate (Caldari System)	1,4 AU
16	Reitsato	Stargate (Caldari System)	1,4 AU
16	Pynekastoh	Stargate (Caldari System)	1,4 AU
32880	---	Venture	93 km
`
	d, err := NewDscan("test_ships.yaml")
	if err != nil {
		t.Fatal(err)
	}

	d1 := d.Scan(rawDscan)

	d2 := DscanResult{
		Types: map[string]int{
			"Black Ops": 2,
			"Force Auxiliary": 3,
			"Frigate": 1,
		},
		Hulls: map[string]int{
			"Apostle": 1,
			"Lif": 2,
			"Marshal": 1,
			"Redeemer": 1,
			"Venture": 1,
		},
		Names: []types.ShipName{
			{
				Ship: "Apostle",
				Name: "Feedot",
			},
			{
				Ship: "Lif",
				Name: "Needs Fit",
			},
			{
				Ship: "Lif",
				Name: "LIF",
			},
			{
				Ship: "Marshal",
				Name: "reddog",
			},
			{
				Ship: "Redeemer",
				Name: "救世啊",
			},
			{
				Ship: "Venture",
				Name: "---",
			},
		},
	}
		
	if ok := cmp.Equal(d1.Hulls, d2.Hulls); !ok {
		t.Error("Hulls comparation fails")
	}
	if ok := cmp.Equal(d1.Types, d2.Types); !ok {
		t.Error("Types comparation fails")
	}
	if ok := cmp.Equal(d1.Names, d2.Names); !ok {
		t.Error("Names compataion fails")
	}
}

