import { h, Component } from 'preact';

export default class Navbar extends Component {
  render() {
    const { modalOpen } = this.props;

    return (
      <div class="navbar has-background-grey-lighter">
        <div class="container">
          <div class="navbar-brand">
            <a href="/" class="navbar-item">
              <img src="/assets/lacks.png" />
            </a>

            <div class="navbar-burger" onClick={modalOpen}>
              <span />
              <span />
              <span />
            </div>
          </div>

          <div class="navbar-menu">
            <div class="navbar-start" />
            <div class="navbar-end">
              <div class="navbar-item is-hoverable">
                <button
                  type="button"
                  class="button is-primary"
                  onClick={modalOpen}
                >
                  Settings
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    );
  }
}
