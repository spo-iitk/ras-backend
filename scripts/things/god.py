import secrets
APC = [
    {
        "name": "Kashishpreet Kaur",
        "email": "ac.2022.kashishpreet@gmail.com"
    },
    {
        "name": "Aditi Phogat",
        "email": "ac.2022.aditi@gmail.com"
    },
    {
        "name": "Afraz Jamal",
        "email": "ac.2022.afraz@gmail.com"
    },
    {
        "name": "Agrim Pandey",
        "email": "ac.2022.agrim@gmail.com"
    },
    {
        "name": "Ananya Agrawal",
        "email": "ac.2022.ananya@gmail.com"
    },
    {
        "name": "Ashutosh Sharma",
        "email": "ac.2022.ashutosh@gmail.com"
    },
    {
        "name": "Kavya Jalan",
        "email": "ac.2022.kavya@gmail.com"
    },
    {
        "name": "Khushbu Kumawat",
        "email": "ac.2022.khushbu@gmail.com"
    },
    {
        "name": "Khushi Gautam",
        "email": "ac.2022.khushi@gmail.com"
    },
    {
        "name": "Nitya Aggarwal",
        "email": "ac.2022.nitya@gmail.com"
    },
    {
        "name": "Payal Singh",
        "email": "ac.2022.payal@gmail.com"
    },
    {
        "name": "Pulkit Dhamija",
        "email": "ac.2022.pulkit@gmail.com"
    },
    {
        "name": "Ravi Patel",
        "email": "ac.2022.ravi@gmail.com"
    },
    {
        "name": "Riktesh Singh",
        "email": "ac.2022.riktesh@gmail.com"
    },
    {
        "name": "Rishabh Yadav",
        "email": "ac.2022.rishabh@gmail.com"
    },
    {
        "name": "Rishi Malhotra",
        "email": "ac.2022.rishi@gmail.com"
    },
    {
        "name": "Sathwika",
        "email": "ac.2022.sathwika@gmail.com"
    },
    {
        "name": "Shivangi Singh",
        "email": "ac.2022.shivangi@gmail.com"
    },
    {
        "name": "Suraj Kumawat",
        "email": "ac.2022.suraj@gmail.com"
    },
    {
        "name": "Upen Mishra",
        "email": "ac.2022.upen@gmail.com"
    },
    {
        "name": "Vandana Basrani",
        "email": "ac.2022.vandana@gmail.com"
    }
]

OPC = [
    {
        "name": "Abhinav D Singh",
        "email": "opc22.abhids@spo.iitk",
    },
    {
        "name": "Sunay Chhajed",
        "email": "opc22.sunay@spo.iitk",
    },
    {
        "name": "Abhishek Kumar",
        "email": "opc22.krabhishek20@spo.iitk",
    },
    {
        "name": "Pragati Singh",
        "email": "opc22.spragati@spo.iitk",
    },
    {
        "name": "Vishwaraj Singh",
        "email": "opc22.vrsingh@spo.iitk",
    },
    {
        "name": "Akhila Mudupu",
        "email": "opc22.akhilam21@spo.iitk",
    }
]

GOD = [
    {
        "name": "Harshit Raj",
        "email": "god22.harshit@spo.iitk"
    },
    {
        "name": "Abhishek Shree",
        "email": "god22.abhishek@spo.iitk"
    },
]

WEB = [
    {
        "name": "Utkarsh Mishra",
        "email": "web22.utkarsh@spo.iitk"
    },
    {
        "name": "Tejas Ahuja",
        "email": "web22.tejas@spo.iitk"
    },
    {
        "name": "Aditya Bangar",
        "email": "web22.aditya@spo.iitk"
    },
    {
        "name": "Krishnansh Agarwal",
        "email": "web22.krish@spo.iitk"
    }
]


# enum
MODE = "SIGNUP"
# MODE = "RESET"

TOKEN = ""

GROUP = OPC

f = open("god.sh", "w")
f.write("#!/bin/bash\n\n")

if MODE == "SIGNUP":
    for x in GROUP:
        f.write("# signup\n")
        f.write("curl 'https://placement.iitk.ac.in/api/auth/god/signup' \\\n")
        f.write("  -H 'Content-Type: application/json' \\\n")
        f.write("  -H 'Authorization: Bearer "+TOKEN+"' \\\n")
        f.write("  -X POST \\\n")
        f.write("  -d '{ \"user_id\": \"" + x["email"] + "\", \"password\": \"" +
                secrets.token_urlsafe(8) + "\",  \"role_id\": 102" + ", \"name\" : \"" + x["name"] + "\" }'\n")
        f.write("\n")
if MODE == "RESET":
    for x in GROUP:
        f.write("# reset\n")
        f.write("curl 'https://placement.iitk.ac.in/api/auth/god/reset-password' \\\n")
        f.write("  -H 'Content-Type: application/json' \\\n")
        f.write("  -H 'Authorization: Bearer "+TOKEN+"' \\\n")
        f.write("  -X POST \\\n")
        f.write("  -d '{ \"user_id\": \"" + x["email"] + "\", \"new_password\": \"" +
                secrets.token_urlsafe(8) + "\" }'\n")
        f.write("\n")
