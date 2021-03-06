import { h, Component } from 'preact';
import Downshift from 'downshift/preact';

export default class SelectMembers extends Component {
  render() {
    // TODO:
    // 1. multiple select...
    // 2. something wrong match logic
    const { members } = this.props;

    return (
      <section class="section">
        <h2 class="title">Members</h2>
        <div class="field">
          <div class="control has-icons-left">
            <Downshift
              onChange={selection => alert(`You selected ${selection.value}`)}
              itemToString={item => (item ? item.value : '')}
            >
              {({
                getInputProps,
                getItemProps,
                isOpen,
                inputValue,
                highlightedIndex,
                selectedItem,
              }) => (
                <div>
                  <input {...getInputProps()} class="input" />
                  <span class="icon is-small is-left">
                    <i class="material-icons">alternate_email</i>
                  </span>
                  <div class="tags">
                    {isOpen
                      ? members
                          .filter(
                            item =>
                              !inputValue || item.value.indexOf(inputValue) > -1
                          )
                          .map((item, index) => {
                            const matched = highlightedIndex === index;
                            const status = matched ? 'is-success ' : 'is-white';
                            return (
                              <span
                                {...getItemProps({
                                  key: item.value,
                                  index,
                                  item,
                                  class: `tag is-medium ${status}`,
                                })}
                              >
                                {item.value}
                              </span>
                            );
                          })
                      : null}
                  </div>
                </div>
              )}
            </Downshift>
          </div>
        </div>
      </section>
    );
  }
}
