import React from 'react'
import AddMemberForm from '../components/AddMemberForm.tsx'

const PoolPage = () => {
    return (
        <>
            <h2>Pool</h2> 
            <AddMemberForm language={"en"} admin={false} uid={"a93140a8-6007-4c93-9a1a-8ede00df9252"}/>
        </>
    )
}

export default PoolPage
