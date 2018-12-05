import { h, Component } from 'preact';
import mdl from 'material-design-lite/material';
import { Button } from 'preact-mdl';

export default class MyButton extends Component {
  render() {
    return(
      <div>
        <Button>I am button!</Button>
      </div>
    )
  }
}
