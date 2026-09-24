import { calculate } from "../../api/calculator";

export interface CalculatorState {
  result: string;
  expression: string;
}

export function calculatorEmpty(): CalculatorState {
  return {
    result: " ",
    expression: "0",
  };
}

export function calculatorAppend(
  state: CalculatorState,
  value: string,
): CalculatorState {
  const expression =
    state.expression === "0" ? value : state.expression + value;
  return { ...state, expression };
}

export function calculatorErase(state: CalculatorState): CalculatorState {
  const expression =
    state.expression.length > 1
      ? state.expression.slice(0, state.expression.length - 1)
      : "0";
  return { ...state, expression };
}

export async function calculatorUpdate(
  state: CalculatorState,
): Promise<CalculatorState> {
  return calculate(state.expression).then((result: string | null) => ({
    ...state,
    result: result === null ? state.result : result,
  }));
}

export function calculatorUseResult(state: CalculatorState): CalculatorState {
  if (isNaN(+state.result)) {
    return state;
  }
  return { ...calculatorEmpty(), expression: state.result };
}
