import React from 'react'
import { useRef, useState } from 'react'
import { languages } from '../assets/languages.ts'
import { postAddMember } from '../apiInterface.ts'

type addMemberFormProps = {
    language: string,
    admin: boolean,
    uid: string,
}

const AddMemberForm = (props: addMemberFormProps) =>  {
    const lang = useRef({})
    const [ name, setName ] = useState("")
    const [ surName, setSurName ] = useState("")
    const [ email, setEmail ] = useState("")
    const [ errorMessage, setErrorMessage ] = useState("")

    if (props.language === "en") {
        lang.current = languages.en
    }
    else if (props.language === "sv") {
        lang.current = languages.sv
    }

    const addUser = async () => {
        if (name === "" || surName === "" || email === "") {
            setErrorMessage(lang.current.errorFields)
            return
        }

        const result = await postAddMember(props.uid, name, surName, email)
        if (!result) {
            setErrorMessage(lang.current.error)
            return
        }
        setErrorMessage("")
    }

    return (
        <>
            <form>
                <label>
                    {lang.current.name}: <br/>
                    <input
                        type="text"
                        value={name}
                        onChange={ (e) => { setName(e.target.value) } }
                    />
                    <br/>{lang.current.surname}: <br/>
                    <input
                        type="text"
                        value={surName}
                        onChange={ (e) => { setSurName(e.target.value) } }
                    />
                    <br/>{lang.current.email}: <br/>
                    <input
                        type="text"
                        value={email}
                        onChange={ (e) => { setEmail(e.target.value) } }
                    />
                </label>
            </form>
            <button onClick={addUser}>{lang.current.enterPool}</button>
            <p className="error">{errorMessage}</p>
        </>
    )
}

export default AddMemberForm

