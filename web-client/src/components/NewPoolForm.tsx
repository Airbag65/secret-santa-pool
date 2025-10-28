import React from "react";
import { useState } from 'react';
import { postCreateNewPool } from '../apiInterface.ts'

const NewPoolForm = () => {

    const [ poolName, setPoolName ] = useState("");
    const [ poolLanguage, setPoolLanguage ] = useState("en");
    const [ poolCurrency, setPoolCurrency ] = useState("");
    const [ poolAmount, setPoolAmount ] = useState(0);
    const [ errorMessage, setErrorMessage ] = useState("")

    const setName = (e) => { setPoolName(e.target.value) }

    const setLanguage = (e) => { setPoolLanguage(e.target.value) }

    const setCurrency = (e) => { setPoolCurrency(e.target.value) }

    const setAmount = (e) => { setPoolAmount(e.target.value) }

    const createPool = async () => {
        const couldCreatePool = await postCreateNewPool(poolName, poolLanguage, poolCurrency, poolAmount)
        if (!couldCreatePool) { setErrorMessage("Something went wrong!") }
    }

    return (
        <>
            <form>
                <label>
                    Pool name: <br/>
                    <input 
                        type="text" 
                        value={poolName} 
                        onChange={setName}
                    />
                    <br/>Language: <br/>
                    <select 
                        value={poolLanguage}
                        onChange={setLanguage}
                    >
                        <option value='en'>English</option>
                        <option value='sv'>Svenska</option>
                    </select>
                    <br/>Currency (eg. SEK for swedish kr, EUR for euro, etc.): <br/>
                    <input 
                        type="text" 
                        value={poolCurrency} 
                        onChange={setCurrency}
                    />
                    <br/>Amount: <br/>
                    <input 
                        type="number" 
                        value={poolAmount} 
                        onChange={setAmount}
                    />
                </label>
            </form>
            <button onClick={createPool}>Create Pool</button>
            <p>{errorMessage}</p>
        </>
    )
}

export default NewPoolForm
