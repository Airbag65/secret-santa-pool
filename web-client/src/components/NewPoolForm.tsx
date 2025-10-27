import React from "react";
import { useState } from 'react';

const NewPoolForm = () => {

    const [ poolName, setPoolName ] = useState("");
    const [ poolLanguage, setPoolLanguage ] = useState("en");
    const [ poolCurrency, setPoolCurrency ] = useState("");
    const [ poolAmount, setPoolAmount ] = useState(0);


    const setName = (e) => {
        setPoolName(e.target.value)
    }

    const setLanguage = (e) => {
        setPoolLanguage(e.target.value)
    }

    const setCurrency = (e) => {
        setPoolCurrency(e.target.value)
    }

    const setAmount = (e) => {
        setPoolAmount(e.target.value)
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
            <p>Current name: {poolName}</p>
            <p>Current language: {poolLanguage}</p>
            <p>Current currency: {poolCurrency}</p>
            <p>Current amount: {poolAmount}</p>
        </>
    )
}

export default NewPoolForm
