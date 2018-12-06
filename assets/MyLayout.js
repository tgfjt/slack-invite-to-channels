import { h, Component } from 'preact';
import { bind } from 'decko';
import Navbar from './Navbar';
import MyForm from './MyForm';

export default class MyLayout extends Component {
  constructor(props) {
    super(props);

    this.state = {
      isModalOpen: false,
      token: '',
      members: [],
      channels: [],
    };
  }

  componentDidMount() {
    const token = localStorage.getItem('__lacks_token__');

    if (token) {
      this.setState({ token });
    }
  }

  componentDidUpdate() {
    const { token, members } = this.state;

    if (token.length < 74 || members.length > 0) return;

    fetch('/members', {
      method: 'post',
      body: JSON.stringify({ token }),
    })
      .then(r => r.json())
      .then(r => {
        if (r.data.ok) {
          this.setState({ members: r.data.members });
        }
      });

    fetch('/channels', {
      method: 'post',
      body: JSON.stringify({ token }),
    })
      .then(r => r.json())
      .then(r => {
        if (r.data.ok) {
          this.setState({ channels: r.data.channels });
        }
      });
  }

  @bind
  modalOpen() {
    this.setState({
      isModalOpen: true,
    });
  }
  @bind
  modalClose() {
    this.setState({
      isModalOpen: false,
    });
  }

  @bind
  saveToken(token) {
    localStorage.setItem('__lacks_token__', this.state.token);
    this.modalClose();
  }

  @bind
  updateTokenValue(e) {
    this.setState({ token: e.target.value.trim() });
  }

  render() {
    const { isModalOpen, token, members, channels } = this.state;
    const modalStyle = isModalOpen ? 'is-active' : '';

    const content = token.length < 74 ? (
      <div class="section">
        <article class="message is-medium">
          <div class="message-header">
            <p>Welcome!</p>
          </div>
          <div class="message-body">
            こんにちは！<br />
            Slack Token を設定してください。<br />
            <p class="has-text-right">
              <button
                  class="button"
                  onClick={this.modalOpen}
                >設定する</button>
            </p>
          </div>
        </article>
      </div>
    ) : (<MyForm token={token} members={members} channels={channels} />)

    return (
      <div>
        <Navbar modalOpen={this.modalOpen} />

        <div class="container">
          {content}
        </div>

        <div class={`modal ${modalStyle}`}>
          <div class="modal-background" />
          <div class="modal-card">
            <header class="modal-card-head">
              <p class="modal-card-title">Settings</p>
              <button
                class="delete"
                aria-label="close"
                onClick={this.modalClose}
              />
            </header>
            <section class="modal-card-body">
              <div class="content">
                <h2>Slack TOKEN</h2>
                <p>
                  To get Slack Token:{' '}
                  <a href="https://api.slack.com/custom-integrations/legacy-tokens">
                    https://api.slack.com/custom-integrations/legacy-tokens
                  </a>
                </p>
                <div class="field">
                  <div class="control">
                    <input
                      class="input"
                      type="text"
                      placeholder="xoxp-xxxx..."
                      value={token}
                      onInput={this.updateTokenValue}
                    />
                  </div>
                </div>
              </div>
            </section>
            <footer class="modal-card-foot">
              <button class="button is-success" onClick={this.saveToken}>
                Save changes
              </button>
              <button class="button" onClick={this.modalClose}>
                Cancel
              </button>
            </footer>
          </div>
        </div>
      </div>
    );
  }
}
