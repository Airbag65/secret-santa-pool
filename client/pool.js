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
    fn.value = ""
    ln.value = ""
    em.value = ""
    try {
        await fetch('http://72.60.16.188:8080/addmember', {
            method: "POST",
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(reqObj)
        })
    } catch {}
    finally {
        location.reload()
    }
} 
copyToClipboard = () => {
  // Get the text field
  // var copyText = document.getElementById("myInput");

  // Select the text field
  // copyText.select();
  // copyText.setSelectionRange(0, 99999); // For mobile devices

   // Copy the text inside the text field
  // navigator.clipboard.writeText(copyText.value);
    navigator.clipboard.writeText(window.location.href.substring(0, window.location.href.length - 11))

  // Alert the copied text
  alert("Copied URL to clipboard: ");
} 

const getPool = async () => {
    if (params.admin == 'true'){
        admin.innerHTML += `<code>${window.location.href.substring(0, window.location.href.length - 11)} </code><button type='button' onclick='copyToClipboard()'>Save to Clipboard</button>`
        admin.innerHTML += "<br><button type='button' onclick='doLottery()'>Perform Lottery</button>"
    }
    const res = await fetch(`http://72.60.16.188:8080/get?uid=${params.uid}`)
    let pool = await res.json().then(element => {return element})
    pool.members.forEach(element => {
        memberList.innerHTML += `<p>${element.first_name} ${element.last_name} | ${element.email}</p>`
    });
}

const doLottery = async () => {
    await fetch(`http://72.60.16.188/lottery?uid=${params.uid}`)
    window.location.href = "./done.html"
}

getPool()
