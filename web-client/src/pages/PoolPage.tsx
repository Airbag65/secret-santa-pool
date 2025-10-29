import React from 'react'
import { useState } from 'react'
import { useSearchParams } from 'react-router'
import AddMemberForm from '../components/AddMemberForm.tsx'
import PoolMembers from '../components/PoolMembers.tsx'
import { getPool } from '../apiInterface.ts'

const PoolPage = () => {
    interface Person {
        email: string;
        first_name: string;
        last_name: string;
    };

    const [ members, setMembers ] = useState<Person[]>([])
    const [ searchParams, setSearchParams ] = useSearchParams()
    const [ loading, setLoading ] = useState(true)

    // useEffect(() => {
    //     setMembers(getPool(searchParams.get('uid')).then((m) => {
    //         return m.members
    //     })) 
    // }, [searchParams])
    if (loading) {
        getPool(searchParams.get('uid')).then((m) => {
            setMembers(m.members)
        }).finally(() => { setLoading(false)} )
    }

    return (
        <>
            <h2>Pool</h2> 
            <PoolMembers language={"en"} members={members} isAdmin={ searchParams.get('admin') == 'true' ? true : false } />
            <AddMemberForm language={"en"} uid={searchParams.get('uid')}/>
        </>
    )
}

export default PoolPage
