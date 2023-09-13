const { createInvoker, getEnvironment } = require("../../../helper.js")
const { ruleId, ruleFile, testBase } = getEnvironment(__dirname)

describe(ruleId, () => {
  const invoke = createInvoker(ruleId, ruleFile, testBase)
  

  test("datatype_in_bulk", () => {
    const testCase = "datatype_in_bulk.rb"
    expect(invoke(testCase)).toMatchSnapshot();
  })
  

  test("datatype_in_index", () => {
    const testCase = "datatype_in_index.rb"
    expect(invoke(testCase)).toMatchSnapshot();
  })
  

  test("datatype_in_update", () => {
    const testCase = "datatype_in_update.rb"
    expect(invoke(testCase)).toMatchSnapshot();
  })
  

  test("ok_datatype_ignored", () => {
    const testCase = "ok_datatype_ignored.rb"
    expect(invoke(testCase)).toMatchSnapshot();
  })
  

  test("ok_no_datatype", () => {
    const testCase = "ok_no_datatype.rb"
    expect(invoke(testCase)).toMatchSnapshot();
  })
  
})