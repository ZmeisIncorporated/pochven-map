package dscan

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ZmeisIncorporated/pochven-map/pkg/types"
	"gopkg.in/yaml.v3"
)


type DscanResult struct {
	Hulls map[string]int
	Types map[string]int
	Names []types.ShipName
}


type Dscan struct {
	Ships types.Ship
}


func NewDscan(path string) (*Dscan, error) {
	dscan := Dscan{
		Ships: types.Ship{},
	}

	fp, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Dscan creating error: %w", err)
	}

	d, err := io.ReadAll(fp)
	if err != nil {
		return nil, fmt.Errorf("%v read error: %w", path, err)
	}
	
	err = yaml.Unmarshal(d, &dscan.Ships)
	if err != nil {
		return nil, fmt.Errorf("%v unmarshal error: %w", path, err)
	}

	return &dscan, nil
}


func (d *Dscan) Scan(rawDscan string) *DscanResult {

	dscanResult := &DscanResult{
		Hulls: make(map[string]int),
		Types: make(map[string]int),
		Names: make([]types.ShipName, 0),
	}
	
	sDscan := strings.Split(rawDscan, "\n")
	for _, v := range sDscan {
		oneLine := strings.Split(v, "\t")
		if len(oneLine) != 4 {
			continue
		}

		shipName := oneLine[1]
		shipHull := oneLine[2]

		shipType, ok := d.Ships[shipHull]
		if !ok {
			// skip not ship
			continue
		}

		dscanResult.Types[shipType] += 1
		dscanResult.Hulls[shipHull] += 1
		dscanResult.Names = append(dscanResult.Names,
			types.ShipName{
				Ship: shipHull,
				Name: shipName,
			})
	}

	return dscanResult
}

