const fs = require('fs')

const filename = process.argv[2];
let f = fs.readFileSync(filename).toString();
let fj = JSON.parse(f)
fj.features.forEach(ft => {
    ft.properties.description = ft.properties.description
        .replaceAll(/<[^<>]*>/g, '')
})
f = JSON.stringify(fj, null, " ")
fs.writeFileSync(filename, f)
