# frozen_string_literal: true

module Buttress
  class AccountNumberMasker
    ALLOWED_NUMERIC_PATTERNS = [
      /\b[[:digit:]]-[[:digit:]]{3}-[[:digit:]]{3}-[[:digit:]]{4}\b/,
      /\b[[:digit:]]{4}-[[:digit:]]{3}-[[:digit:]]{4}\b/,
      /\b[[:digit:]]{3}-[[:digit:]]{3}-[[:digit:]]{4}\b/,
      /\b[[:digit:]]{4}-[[:digit:]]{7}\b/,
      /\b[[:digit:]]{3}-[[:digit:]]{7}\b/,
      /\b[[:digit:]]{3}-[[:digit:]]{4}\b/
    ].freeze

    ##
    # Masks account numbers.
    # Does not mask anything that matches an allowed pattern.
    #
    def self.mask(desc)
      self.mask!(desc.dup)
    end

    def self.mask!(desc)
      return desc if desc.blank?
      return desc if /(CHECK|DRAFT|SHARE DRAFT|DFT DEBIT)[[[:space:]]]*[#]*[[[:space:]]]*[[[:digit:]]]{1,}/i =~ desc

      # Looks for 7 or more digits and/or dashes
      masked_description = desc.gsub(/[-[[:digit:]]]{7,}/) do |num|
        puts num
        if ALLOWED_NUMERIC_PATTERNS.none? { |pattern| num =~ pattern }
          num.gsub!(/[[:digit:]](?=[-[[:digit:]]]{4})/, "X")
        end

        num
      end

      masked_description
    end
  end
end

class String
  def blank?
    false
  end
end

# "555-5555 5555-5555-5555-5555" => "555-5555 XXXX-XXXX-XXXX-5555"

puts ::Buttress::AccountNumberMasker.mask!("555-5555 5555-5555-5555-5555")
