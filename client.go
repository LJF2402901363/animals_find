/*
@author: 陌意随影
@since: 2023/11/6 00:39:08
@desc:
*/
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type GetDateRes struct {
	Success  bool `json:"success"`
	Response struct {
		Result struct {
			AtomicalId                   string      `json:"atomical_id"`
			TopLevelRealmAtomicalId      string      `json:"top_level_realm_atomical_id"`
			TopLevelRealmName            string      `json:"top_level_realm_name"`
			NearestParentRealmAtomicalId string      `json:"nearest_parent_realm_atomical_id"`
			NearestParentRealmName       string      `json:"nearest_parent_realm_name"`
			RequestFullRealmName         string      `json:"request_full_realm_name"`
			FoundFullRealmName           string      `json:"found_full_realm_name"`
			MissingNameParts             interface{} `json:"missing_name_parts"`
			Candidates                   []struct {
				TxNum                int    `json:"tx_num"`
				AtomicalId           string `json:"atomical_id"`
				Txid                 string `json:"txid"`
				CommitHeight         int    `json:"commit_height"`
				RevealLocationHeight int    `json:"reveal_location_height"`
			} `json:"candidates"`
			NearestParentRealmSubrealmMintRules struct {
				NearestParentRealmAtomicalId string      `json:"nearest_parent_realm_atomical_id"`
				CurrentHeight                int         `json:"current_height"`
				CurrentHeightRules           interface{} `json:"current_height_rules"`
			} `json:"nearest_parent_realm_subrealm_mint_rules"`
			NearestParentRealmSubrealmMintAllowed bool `json:"nearest_parent_realm_subrealm_mint_allowed"`
		} `json:"result"`
	} `json:"response"`
}

type GetDataReq struct {
	Params []interface{} `json:"params"`
}

func IsNotFind(params []interface{}) (bool, error) {
	apiURL := "https://ep.atomicals.xyz/proxy/blockchain.atomicals.get_realm_info"
	marshal, err := json.Marshal(GetDataReq{Params: params})
	if err != nil {
		return false, err
	}
	reader := strings.NewReader(string(marshal))
	request, err := http.NewRequest("POST", apiURL, reader)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Origin", "https://www.satsx.io")
	request.Header.Add("Referer", "https://www.satsx.io/")
	request.Header.Add("authority", "ep.atomicals.xyz")
	request.Header.Add("method", "ep.atomicals.xyz")
	request.Header.Add("scheme", "https")
	request.Header.Add("path", "/proxy/blockchain.atomicals.get_realm_info")
	request.Header.Add("Accept-Encoding", "gzip, deflate, br")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	readAll, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	var res GetDateRes
	err = json.Unmarshal(readAll, &res)
	if err != nil {
		return false, err
	}
	return res.Success && len(res.Response.Result.Candidates) == 0, nil
}

var baseStr = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "s", "y", "z",
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

var randomStrMap = map[string]struct{}{}
var notFindMap = map[string]bool{}

func InitRandomStrMap() {
	for _, i := range baseStr {
		for _, j := range baseStr {
			for _, k := range baseStr {
				randomStr := fmt.Sprintf("%s%s%s", i, j, k)
				randomStrMap[randomStr] = struct{}{}
			}
		}
	}
}

func CreateReq() {
	if len(randomStrMap) == 0 {
		InitRandomStrMap()
	}
	for key := range randomStrMap {
		param := []interface{}{key, 0}
		find, err := IsNotFind(param)
		if err != nil {
			continue
		}
		if !find {
			continue
		}

		file, err := os.OpenFile("sub_data.txt", os.O_APPEND|os.O_WRONLY, os.ModePerm)
		if err != nil {
			return
		}
		_, err = file.WriteString(fmt.Sprintf("%s\n", key))
		if err != nil {
			return
		}
		notFindMap[key] = find
	}
	marshal, err := json.Marshal(notFindMap)
	if err != nil {
		return
	}
	err = os.WriteFile("total_data.json", marshal, os.ModePerm)
	if err != nil {
		return
	}
}
