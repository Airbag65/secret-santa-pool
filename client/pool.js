const memberList = document.getElementById('memberList')
const admin = document.getElementById('admin')
let paramsList = window.location.href.split('?')[1]
let params = {}
params.admin = 'false'
paramsList = paramsList.split('&');
paramsList.forEach(element => {
    let param = element.split('=');
    params[param[0]] = param[1] 
});
// console.log(params)

const addMember = async () => {
    const fn = document.getElementById("first_name")
    const ln = document.getElementById("last_name")
    const em = document.getElementById("email")
    reqObj = {
        uid: params.uid,
        person: {
            first_name: fn.value,
            last_name: ln.value,
            email: em.value
        }
    }
    try {
        await fetch('http://127.0.0.1:8080/addmember', {
            method: "POST",
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(reqObj)
        })
    } catch {}
    // console.log(reqObj)
} 

const getPool = async () => {
    if (params.admin == 'true'){
        admin.innerHTML += "does this work?"
    }
    const res = await fetch(`http://127.0.0.1:8080/get?uid=${params.uid}`)
    let pool = await res.json().then(element => {return element})
    // console.log(pool.members)
    pool.members.forEach(element => {
        memberList.innerHTML += `<p>${element.first_name} ${element.last_name} | ${element.email}</p>`
    });
}

getPool()
