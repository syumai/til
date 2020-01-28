import * as React from 'react';
const { useState } = React;

interface EchoFormSubmitHandler {
  (body: string): Promise<void>;
}

interface EchoFormParams {
  title: string;
  onSubmit: EchoFormSubmitHandler;
}

const EchoForm = ({ title, onSubmit }: EchoFormParams) => {
  const [body, setBody] = useState('');

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    await onSubmit(body);
    setBody('');
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setBody(e.target.value);
  };

  return (
    <>
      <h2>{title}</h2>
      <form onSubmit={handleSubmit}>
        <input type="text" value={body} onChange={handleChange}></input>
        <button>Submit</button>
      </form>
    </>
  );
};

export const App = () => {
  const onSubmitEcho: EchoFormSubmitHandler = async body => {
    console.log('onSubmitEchoForm', body);
  };
  const onSubmitEchoReverse: EchoFormSubmitHandler = async body => {
    console.log('onSubmitEchoReverseForm', body);
  };
  return (
    <div>
      <h1>gRPC Web Echo Server Example</h1>
      <EchoForm title="Echo Form" onSubmit={onSubmitEcho} />
      <EchoForm title="Echo Reverse Form" onSubmit={onSubmitEchoReverse} />
    </div>
  );
};
