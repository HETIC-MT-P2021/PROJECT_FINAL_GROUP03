import Form from "react-bootstrap/Form";

function WelcomeMessageForm() {
  return (
    <Form>
      <Form.Group className="mb-3" controlId="formBasicEmail">
        <Form.Label>Welcome message</Form.Label>
        <Form.Control type="text" placeholder="Welcome here!" />
      </Form.Group>
    </Form>
  );
}

export default WelcomeMessageForm;
