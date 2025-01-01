const createPool = async () => {
    try {
        // const res = await fetch("http://127.0.0.1:8080")
        // console.log(res.ok)
        const res = await fetch("http://127.0.0.1:8080/createpool", {
            method: 'POST'
        }) 
        if(!res.ok){
            console.log(res.status)
        }
        const json = await res.json() 
        console.log(json)

    } catch (error) {
        console.log(error)
    }
}
