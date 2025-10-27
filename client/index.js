const createPool = async () => {
    try {
        const res = await fetch("http://72.60.16.188:8080/createpool", {
            method: 'POST'
        }) 
        if(!res.ok){
            console.log(res.status)
            return
        }         
        const json = await res.json() 
        console.log(json.uuid)
        window.location.href = `./pool?uid=${json.uuid}&admin=true`

    } catch (error) {
        console.log(error)
        return
    }
}
