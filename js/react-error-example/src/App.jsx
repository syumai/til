import React from 'react';

class ErrorBoundary extends React.Component {
  constructor(props) {
    super(props);
    this.state = { hasError: false };
  }

  static getDerivedStateFromError(error) {
    // Update state so the next render will show the fallback UI.
    return { hasError: true };
  }

  componentDidCatch(error, info) {
    console.error("componentDidCatch", error)
  }

  render() {
    if (this.state.hasError) {
      // You can render any custom fallback UI
      return this.props.fallback;
    }

    return this.props.children;
  }
}
let counter = 0;
const C = () => {
  counter++;
  console.log("rendered!", counter)
  if (counter % 2 === 1)
    throw new Error("hoge")
}

const FB = () => {
  console.log("render fallback!", counter)
  return <div>fallback</div>
}

export function App(props) {
  return (
    <div className='App'>
      <ErrorBoundary fallback={<FB />}>
        <C />
        <h1>Hello React.</h1>
        <h2>Start editing to see some magic happen!</h2>
      </ErrorBoundary>
    </div>
  );
}
