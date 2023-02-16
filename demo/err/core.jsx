// const { useState } = React;
function MyApp() {
  const [state, setState] = React.useState({ msg: "暂无消息", details: [] });
  const [view, setView] = React.useState(false);
  return (
    <div>
      <div
        className="ui primary button"
        onClick={async () => {
          fetch("/api", {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
          })
            .then((resp) => {
              resp.json().then((json) => {
                console.log("🚀 ~ file: core.jsx:20 ~ resp.text ~ text", json);
                setState({ msg: json["msg"], details: json["details"] });
              });
            })
            .catch((err) => {
              alert(err);
            });
        }}
      >
        发起请求
      </div>
      <div
        style={{
          marginTop: "30px",
        }}
        className="ui message"
      >
        <div>{state.msg}</div>
        <div className="ui button mini " onClick={() => setView(!view)}>
          {view ? "看不懂，隐藏吧" : "点我查看细节"}
        </div>
        {view && (
          <ul className="ui list">
            {state.details.map((v, idx) => {
              return <li key={idx}>{v}</li>;
            })}
          </ul>
        )}
      </div>
    </div>
  );
}

const container = document.getElementById("root");
// @ts-ignore
const root = ReactDOM.createRoot(container);
root.render(<MyApp />);
