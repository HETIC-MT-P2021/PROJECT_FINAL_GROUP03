import {Button, Form} from "react-bootstrap";
import {MdDelete, MdEdit} from "react-icons/all";
import {ChangeEvent, useState} from "react";

const ForbiddenWordItem = (props: {
    index: number,
    word: string,
    onDelete: (index: number) => void,
    onEdit: (index: number, value: string) => void
}) => {
    const [index, setIndex] = useState(props.index);
    const [word, setWord] = useState(props.word);
    const [isEditMode, setEditMode] = useState(false);

    const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
        setWord(event.currentTarget.value);
    };

    const deleteWord = () => {
        props.onDelete(index);
    };
    let wordTempl = <td>{word}</td>;

    const activateEditMode = () => {
        setEditMode(true);
    }
    const deActivateEditMode = () => {
        props.onEdit(index, word);
        setEditMode(false);
    }

    if(isEditMode) {
        wordTempl = <td><Form.Control
            type="text"
            value={word}
            placeholder="basic forbidden word"
            onChange={handleChange}
            onBlur={deActivateEditMode}
        /></td>
    }

    return (
      <tr key={index}>
          { wordTempl }
          <td>
              <Button className="td-action" variant="primary" onClick={activateEditMode} ><MdEdit/></Button>
              <Button className="td-action" variant="danger" onClick={deleteWord}><MdDelete/></Button>
          </td>
      </tr>
  );
};

export default ForbiddenWordItem;