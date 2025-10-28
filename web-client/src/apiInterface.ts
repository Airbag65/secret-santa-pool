// import "dotenv/config"
// dotenv.config({path: '../.env'})


interface newPoolReq {
    language: string;
    currency: string;
    pool_name: string;
    amount: number;
}

const postCreateNewPool: boolean = async (name: string, language: string, currency: string, amount: number) => {
    const IP_ADDRESS = import.meta.env.VITE_ENV_IP

    const requestBody: newPoolReq = {
        language: language,
        currency: currency,
        pool_name: name,
        amount: Number(amount)
    }

    console.log(`http://${IP_ADDRESS}:8080/createpool`)
    
    console.log(JSON.stringify(requestBody))

    const response = await fetch(`http://${IP_ADDRESS}:8080/createpool`, 
        {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(requestBody)
        })
    console.log(response)
    
    if (!response.ok) { return false }
    
    const resData = await response.json()

    window.location.href = `./pool?uid=${resData.uuid}`

    return true
}

export { postCreateNewPool }
