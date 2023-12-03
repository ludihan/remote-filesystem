package html

var Index = `
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <title>Access</title>
        <script src="https://unpkg.com/htmx.org@1.9.9" integrity="sha384-QFjmbokDn2DjBjq+fM+8LUIVrAgqcNW2s0PjAxHETgRn9l4fvX31ZxDxvwQnyMOX" crossorigin="anonymous"></script>
    </head>
    <body>
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
        </form>
    </body>
</html>
`

var Form = `
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
        </form>
`
