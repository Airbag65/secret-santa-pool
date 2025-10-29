interface newPoolReq {
    language: string;
    currency: string;
    pool_name: string;
    amount: number;
}

const IP_ADDRESS = import.meta.env.VITE_ENV_IP

const postCreateNewPool: boolean = async (name: string, language: string, currency: string, amount: number) => {

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


interface addMemberReq {
    uid: string;
    person: Person;
}

interface Person {
    email: string;
    first_name: string;
    last_name: string;
}

const postAddMember: boolean = async (uid: string, name: string, surname: string, email: string) => {
    const newMember = {
        email: email,
        first_name: name,
        last_name: surname
    }

    const requestBody: addMemberReq = {
        uid: uid,
        person: newMember
    }

    const response = await fetch(`http://${IP_ADDRESS}:8080/addmember`,
        {
            method: 'POST',
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify(requestBody)
        })
    if (!response.ok) { return false }
    
    return true
}

const getPool: Person[] = async (uid: string) => {
    const response = await fetch(`http://${IP_ADDRESS}:8080/get?uid=${uid}`)
    if (!response.ok) { return [] }
    return await response.json() as Person[]
}

export { postCreateNewPool, postAddMember, getPool }
