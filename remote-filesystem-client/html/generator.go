package html

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
