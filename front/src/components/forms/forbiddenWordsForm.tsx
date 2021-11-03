import {useState} from "react";

const ForbiddenWordsForm = (props: {
    forbidden_words: string;
    onvalidate: (wordsList: string) => void;
}) => {
    const [forbiddenWords, setForbiddenWords] = useState(props.forbidden_words);
    let arrayWords = forbiddenWords.split(",");
    const wordsItems = arrayWords.map(
        (word, index) => <li key={index}>{ word }</li>
    );
    const saveData = () => {
        props.onvalidate(forbiddenWords || "");
    };

    return (
        <div>
            <b>Liste de mots interdits</b>
            <ul>{wordsItems}</ul>
        </div>
    )
};

export default ForbiddenWordsForm;