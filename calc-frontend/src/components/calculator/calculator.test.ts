import { calculate } from "../../api/calculator";
import {
  calculatorAppend,
  calculatorEmpty,
  calculatorErase,
  calculatorUpdate,
  calculatorUseResult,
} from "./calculator";
import { expect, test, vi } from "vitest";

vi.mock("../../api/calculator.ts", () => ({
  calculate: vi.fn(),
}));

test("empty calculator", () => {
  expect(calculatorEmpty()).toEqual({ result: " ", expression: "0" });
});

test("calculator append", () => {
  expect(calculatorAppend({ result: " ", expression: "0" }, "1")).toEqual({
    result: " ",
    expression: "1",
  });

  expect(calculatorAppend({ result: " ", expression: "123" }, "1")).toEqual({
    result: " ",
    expression: "1231",
  });

  expect(calculatorAppend({ result: "3", expression: "1+2" }, "+")).toEqual({
    result: "3",
    expression: "1+2+",
  });
});

test("calculator erase", () => {
  expect(calculatorErase({ result: " ", expression: "0" })).toEqual({
    result: " ",
    expression: "0",
  });

  expect(calculatorErase({ result: " ", expression: "123" })).toEqual({
    result: " ",
    expression: "12",
  });

  expect(calculatorErase({ result: "3", expression: "1+2" })).toEqual({
    result: "3",
    expression: "1+",
  });
});

test("calculator update", async () => {
  vi.mocked(calculate).mockResolvedValue("123");
  expect(
    await calculatorUpdate({ result: " ", expression: "100+20+3" }),
  ).toEqual({
    result: "123",
    expression: "100+20+3",
  });
  expect(calculate).toHaveBeenCalledOnce();

  vi.mocked(calculate).mockResolvedValue("0");
  expect(await calculatorUpdate({ result: " ", expression: "0" })).toEqual({
    result: "0",
    expression: "0",
  });
  expect(calculate).toHaveBeenCalledTimes(2);

  vi.mocked(calculate).mockResolvedValue("new kind of error");
  expect(await calculatorUpdate({ result: " ", expression: "1/0" })).toEqual({
    result: "new kind of error",
    expression: "1/0",
  });
  expect(calculate).toHaveBeenCalledTimes(3);
});

test("calculator use result", async () => {
  expect(
    calculatorUseResult({ result: "123", expression: "100+20+3" }),
  ).toEqual({
    result: " ",
    expression: "123",
  });
});

expect(
  calculatorUseResult({ result: "error text", expression: "100+20+3" }),
).toEqual({
  result: "error text",
  expression: "100+20+3",
});
