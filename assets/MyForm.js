import { h, Component } from 'preact';
import SelectMembers from './SelectMembers';
import SelectChannels from './SelectChannels';

export default class MyForm extends Component {
  render() {
    const members = this.props.members.map(m => ({
      value: m.profile.display_name,
    }));
    const channels = this.props.channels.map(m => ({
      value: m.name,
    }));

    return (
      <div class="columns">
        <div class="column is-half">
          <SelectMembers members={members} />
        </div>
        <div class="column is-half">
          <SelectChannels channels={channels} />
        </div>
      </div>
    );
  }
}
