import {useEffect, useState} from "react";
import {Table, Button} from "react-bootstrap";
import "./forbiddenWords.css";
import ForbiddenWordItem from "./forbiddenWordItem";

const ForbiddenWordsForm = (props: {
    forbidden_words: string[];
    onvalidate: (wordsList: string) => void;
}) => {
    const [forbiddenWords, setForbiddenWords] = useState(props.forbidden_words);

    const deleteWord = (index: number) => () => {
        const newArray = [...forbiddenWords];
        newArray.splice(index, 1)
        console.log(newArray)
        setForbiddenWords(newArray);
    };

    const replaceWord = (index: number, value: string) => {
        let newArray = [...forbiddenWords];
        newArray[index] = value;
        setForbiddenWords(newArray);
    };

    const saveData = () => {
        const sanitizedList = forbiddenWords.filter(word => word.length > 1)
        props.onvalidate(sanitizedList.join(",") || "");
        setForbiddenWords(sanitizedList);
    };
    
    const addWord = () => {
        setForbiddenWords(forbiddenWords.concat(""));
    };

    return (
        <div className="forbidden-words">
            <h3>Mots interdits</h3>
            <div className="buttons-container">
                <Button className="buttons-container__button" variant="primary" onClick={addWord}>
                    Nouveau
                </Button>
                <Button className="buttons-container__button" variant="success" onClick={saveData}>
                    Enregistrer
                </Button>
            </div>
            <Table>
                <thead>
                    <tr>
                        <th>Mot</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody>
                {
                    forbiddenWords.map(
                        (word, index) => <ForbiddenWordItem key={word} word={word} index={index} onDelete={deleteWord(index)} onEdit={replaceWord}/>
                    )
                }
                </tbody>
            </Table>
        </div>
    )
};

export default ForbiddenWordsForm;