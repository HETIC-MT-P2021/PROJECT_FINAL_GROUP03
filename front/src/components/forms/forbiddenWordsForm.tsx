import {useState} from "react";
import {Table, Button} from "react-bootstrap";
import "./forbiddenWords.css";
import {MdDelete, MdEdit} from "react-icons/all";

const ForbiddenWordsForm = (props: {
    forbidden_words: string;
    onvalidate: (wordsList: string) => void;
}) => {
    const [forbiddenWords, setForbiddenWords] = useState(props.forbidden_words);
    const [wordsArray, setWordsArray] = useState([""]);
    setWordsArray(forbiddenWords.split(","));

    const wordsItems = wordsArray.map(
        (word, index) => <tr key={index}>
            <td>{ word }</td>
            <td>
                <Button className="td-action" variant="primary"><MdEdit/></Button>
                <Button className="td-action" variant="danger" ><MdDelete/></Button>
            </td>
        </tr>
    );
    const saveData = () => {
        props.onvalidate(forbiddenWords || "");
    };
    const addWord = () => {

    };
    return (
        <div className="forbidden-words">
            <h3>Mots interdits</h3>
            <Table>
                <thead>
                <tr>
                    <th>Mot</th>
                    <th>Actions</th>
                </tr>
                </thead>
                <tbody>
                { wordsItems}
                </tbody>
            </Table>
            <Button variant="success" onClick={saveData}>
                Enregistrer
            </Button>
            <Button variant="success" onClick={addWord}>
                Nouveau
            </Button>
        </div>
    )
};

export default ForbiddenWordsForm;