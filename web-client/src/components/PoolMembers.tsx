import React from 'react'
import languages from '../assets/languages.ts'
import { useRef } from 'react'

interface Person {
    email: string;
    first_name: string;
    last_name: string;
};


type poolMembersProps = {
    language: string,
    members: Person[],
    isAdmin: boolean
}

const PoolMembers = (props: poolMembersProps) => {
    const lang = useRef({})

    if (props.language === 'en') {
        lang.current = languages.en
    }
    else if (props.language === 'sv') {
        lang.current = languages.sv
    }

    console.log(props.members)
    
    return (
        <>
        <p></p>
        {
            props.members.forEach((member) => {return (<p>{member.first_name} {member.last_name} | {member.email}</p>)})
        }
        </>
    )
}

export default PoolMembers
