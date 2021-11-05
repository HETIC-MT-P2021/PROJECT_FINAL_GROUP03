import "./birthdayMessageForm.css"
import {ChangeEvent, useState} from "react";
import {Button, Form} from "react-bootstrap";

const BirthdayMessageForm = (props: {
    birthday_message: string;
    onvalidate: (message: string) => void;
}) => {
    const [birthdayMessage, setBirthdayMessage] = useState(props.birthday_message);

    const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
        setBirthdayMessage(event.currentTarget.value);
    };

    const saveData = () => {
        props.onvalidate(birthdayMessage || "");
    };

    return (
        <Form className="">
            <Form.Group className="mb-3" controlId="formBasicEmail">
                <Form.Label><h3>Message d'anniversaire</h3></Form.Label>
                <Form.Control
                    type="text"
                    value={birthdayMessage}
                    placeholder="Birthday message here!"
                    onChange={handleChange}
                />
                <br />
                <Button variant="success" onClick={saveData}>
                    Enregistrer
                </Button>
            </Form.Group>
        </Form>
    );
}

export default BirthdayMessageForm;