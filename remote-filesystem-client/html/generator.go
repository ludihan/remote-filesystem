package html

import (
    "encoding/json"
    "log"
    "strconv"
)

func ReturnErrorForm(err string) string {
    return `
        <form hx-post="/login" hx-target="this">
            <h1>Access</h1>
            <ul>
                <li>
                    <label for="address">port:</label>
                    <input type="text" id="address" name="user_address" />
                </li>
                <li>
                    <button type="submit">Submit</button>
                </li>
            </ul>
            ` +
            `<p>` + err + `</p>` + `</form>`

}

type currentDir struct {
    Dir   []string `json:"dir"`
	Files []files `json:"files"`
    Home  string   `json:"home"`
	Path  string   `json:"path"`
}

type files struct {
    Last_modification string `json:"last_modification"`
    Name              string `json:"name"`
    Size              int `json:"size"`
}

func ReturnDir(body []byte) string {
    var dir currentDir
    err := json.Unmarshal(body, &dir)
    if err != nil {
        log.Println("oii sou eu de novo")
        log.Fatal(err)
    }
    var resposeString string
    resposeString += `<h1>` + dir.Path + `</h1>`
    resposeString += `<table><tbody>`
    if dir.Path == "" {
        dir.Path = dir.Home
    }

    resposeString += `<tr><th>Name</th><th>Size</th><th>Last Modified</th></tr>`
    for _,element := range dir.Dir {
        resposeString += `<tr><td>` + element + `</tr></td>`
    }
    for _,element := range dir.Files {
        resposeString += `<tr><td>` + element.Name + `</td>`
        resposeString += `<td>` + strconv.Itoa(element.Size) + `</td>`
        resposeString += `<td>` + element.Last_modification + `</td></tr>`
    }
    resposeString += `</tbody></table>`
    log.Println(resposeString)
    return resposeString
}

