import {useEffect, useState} from "react";
import {Table, Button} from "react-bootstrap";
import "./forbiddenWords.css";
import ForbiddenWordItem from "./forbiddenWordItem";

const ForbiddenWordsForm = (props: {
    forbidden_words: string[];
    onvalidate: (wordsList: string) => void;
}) => {
    const [forbiddenWords, setForbiddenWords] = useState(props.forbidden_words);

    const saveData = () => {
        props.onvalidate(forbiddenWords.join(",") || "");
    };

    const addWord = () => {
        setForbiddenWords(forbiddenWords.concat(""));
    };

    const deleteWord = (index: number) => {
        const newArray = forbiddenWords;
        newArray.splice(index, 1)
        setForbiddenWords(newArray);
    };

    const replaceWord = (index: number, value: string) => {
        let newArray = forbiddenWords;
        newArray[index] = value;
        setForbiddenWords(newArray);
    };

    const wordsItems = forbiddenWords.map(
        (word, index) => <ForbiddenWordItem key={index} word={word} index={index} onDelete={deleteWord} onEdit={replaceWord}/>
    );
    
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
            <div className="buttons-container">
                <Button className="buttons-container__button" variant="primary" onClick={addWord}>
                    Nouveau
                </Button>
                <Button className="buttons-container__button" variant="success" onClick={saveData}>
                    Enregistrer
                </Button>
            </div>
        </div>
    )
};

export default ForbiddenWordsForm;