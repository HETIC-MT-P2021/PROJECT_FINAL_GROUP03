import "./welcomeMessageForm.css"
import {ChangeEvent, useState} from "react";
import {Button, Form} from "react-bootstrap";

const WelcomeMessageForm = (props: {
    welcome_message: string | undefined;
    onvalidate: (message: string) => void;
}) => {
    const [welcomeMessage, setWelcomeMessage] = useState(props.welcome_message);

    const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
        setWelcomeMessage(event.currentTarget.value);
    };

    const saveData = () => {
        props.onvalidate(welcomeMessage || "");
    };

    return (
        <Form>
            <Form.Group className="mb-3" controlId="formBasicEmail">
                <Form.Label>Welcome message</Form.Label>
                <Form.Control
                    type="text"
                    value={welcomeMessage}
                    placeholder="Welcome here!"
                    onChange={handleChange}
                />
                <br />
                <Button variant="success" onClick={saveData}>
                    Save
                </Button>{" "}
            </Form.Group>
        </Form>
    );
}

export default WelcomeMessageForm;