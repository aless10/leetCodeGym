use std::collections::HashMap;

impl Solution {
    pub fn contains_duplicate(nums: Vec<i32>) -> bool {
        let mut counter: HashMap<String, i32> = HashMap::new();
        for num in nums {
            let key = num.to_string();
            let count = counter.entry(key.clone()).or_insert(0);
            *count += 1;
            if (*count > 1) {
                return true;
            }
        }
        return false;
        
    }
}