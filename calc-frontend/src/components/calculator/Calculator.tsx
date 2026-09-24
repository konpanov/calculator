import { useState, useCallback } from 'react';
import './Calculator.css';
import { calculatorAppend, calculatorEmpty, calculatorErase, calculatorUpdate, calculatorUseResult, type CalculatorState } from '../calculator/calculator';

export default function Calculator() {
  const [calc, setCalc] = useState<CalculatorState>({ result: ' ', expression: '0' });

  const append = useCallback((value: string) => {
    calculatorUpdate(calculatorAppend(calc, value)).then(setCalc);
  }, [calc]);

  const equals = useCallback(() => {
    calculatorUpdate(calc).then(calculatorUseResult).then(setCalc);
  }, [calc]);

  const clearAll = useCallback(() => {
    setCalc(calculatorEmpty())
  }, []);

  const erase = useCallback(() => {
    calculatorUpdate(calculatorErase(calc)).then(setCalc);
  }, [calc]);

  const keys = [
    { label: 'AC', cls: 'fn', onClick: clearAll },
    { label: '<<', cls: 'fn', onClick: erase },
    { label: '(', cls: `fn`, onClick: () => append('(') },
    { label: ')', cls: `fn`, onClick: () => append(')') },
    { label: '%', cls: 'op', onClick: () => append('%') },
    { label: '√', cls: `op`, onClick: () => append('√') },
    { label: '^', cls: `op`, onClick: () => append('^') },
    { label: '÷', cls: `op`, onClick: () => append('÷') },
    { label: '7', cls: 'num', onClick: () => append('7') },
    { label: '8', cls: 'num', onClick: () => append('8') },
    { label: '9', cls: 'num', onClick: () => append('9') },
    { label: '×', cls: `op`, onClick: () => append('×') },
    { label: '4', cls: 'num', onClick: () => append('4') },
    { label: '5', cls: 'num', onClick: () => append('5') },
    { label: '6', cls: 'num', onClick: () => append('6') },
    { label: '−', cls: `op`, onClick: () => append('−') },
    { label: '1', cls: 'num', onClick: () => append('1') },
    { label: '2', cls: 'num', onClick: () => append('2') },
    { label: '3', cls: 'num', onClick: () => append('3') },
    { label: '+', cls: `op`, onClick: () => append('+') },
    { label: '0', cls: 'num span2', onClick: () => append('0') },
    { label: '.', cls: 'num', onClick: () => append('.') },
    { label: '=', cls: 'eq', onClick: equals },
  ];

  return (
    <div className="calc">
      <div className="display">
        <div className="prev">{calc.result}=</div>
        <div className="curr">{calc.expression}</div>
      </div>
      <div className="keys">
        {keys.map((k, i) => (
          <button key={i} className={`key ${k.cls}`} onClick={k.onClick}>
            {k.label}
          </button>
        ))}
      </div>
    </div>
  );
}
