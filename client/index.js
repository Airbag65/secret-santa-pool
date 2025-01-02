const createPool = async () => {
    try {
        const res = await fetch("http://127.0.0.1:8080/createpool", {
            method: 'POST'
        }) 
        if(!res.ok){
            console.log(res.status)
            return
        }         
        const json = await res.json() 
        console.log(json.uuid)
        window.location.href = `./pool.html?uid=${json.uuid}`

    } catch (error) {
        console.log(error)
        return
    }
}
